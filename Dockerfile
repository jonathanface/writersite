FROM node:19-bullseye AS frontend-builder
ARG VITE_STRIPE_KEY
ARG VITE_MODE
WORKDIR /app
COPY ./ui/package*.json ./
RUN npm install
COPY ./ui/src ./src
COPY ./ui/index.html ./
COPY ./ui/public ./public
COPY ./ui/tsconfig.json ./
COPY ./ui/tsconfig.app.json ./
COPY ./ui/tsconfig.node.json ./
COPY ./ui/vite.config.ts ./
RUN npm run build

FROM golang:1.24 AS backend-builder
# Install wkhtmltox dependencies
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    fontconfig \
    libjpeg62-turbo \
    libx11-6 \
    libxcb1 \
    libxext6 \
    libxrender1 \
    xfonts-75dpi \
    xfonts-base \
    pandoc \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Set environment variables for Go
ENV GO111MODULE=auto \
    GOPATH=/go \
    PATH=$GOPATH/bin:/usr/local/go/bin:/usr/local/bin:/usr/local/:$PATH

ENV PORT=":80"

COPY --from=frontend-builder /app/dist /app/static/rd-ui/dist

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
COPY ./api ./api
COPY ./models ./models
COPY ./main.go ./main.go

RUN go mod tidy

RUN mkdir -p ./tmp
RUN go build -o /jonathanface
#CMD ["sleep", "infinity"]
CMD ["/jonathanface"]


# -------- Frontend build --------
FROM node:22-alpine AS frontend-builder

WORKDIR /app/ui
# Only copy manifest first for better caching
COPY ui/package*.json ./
RUN npm ci

# Copy the rest of the UI
COPY ui/ ./

# Build with explicit mode (defaults to production if not provided)
ARG VITE_MODE=production
ENV NODE_ENV=production
RUN npm run build -- --mode=${VITE_MODE}

# -------- Backend build --------
FROM golang:1.24-alpine AS backend-builder

WORKDIR /src

# Enable modules & static build
ENV CGO_ENABLED=0 GOOS=linux

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy backend source
COPY api ./api
COPY models ./models
COPY main.go ./main.go

# Bring in built UI assets
COPY --from=frontend-builder /app/ui/dist ./ui

# Build optimized binary
RUN go build -ldflags="-s -w" -o /out/app ./main.go

# -------- Minimal runtime --------
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app
# Copy binary and UI
COPY --from=backend-builder /out/app ./app
COPY --from=backend-builder /src/ui ./ui

# Runtime env
ENV PORT=80 \
    MODE=production \
    GIN_MODE=release

EXPOSE 80

CMD ["./app"]

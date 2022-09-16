echo Starting...
while getopts m: flag
do
  case "${flag}" in
    m) method=${OPTARG};;
  esac
done

if [[ "$method" == "all" ]]; then
  echo Linting JavaScript...
  npm --prefix ./ui run checkAndFix
fi
echo Compiling JavaScript...
npm --prefix ./ui run buildUI
if [[ "$method" == "all" ]]; then
  echo Formatting GoLang...
  go fmt ./...
  echo Compiling and executing golang build...
  go build && go run main.go
fi
echo Starting...
while getopts m: flag
do
  case "${flag}" in
    m) method=${OPTARG};;
  esac
done
echo PARAM $method
if [[ "$method" == "all" ]]; then
  echo Linting JavaScript...
  npm --prefix run checkAndFix
fi
echo Compiling JavaScript...
npm --prefix ./ui run buildUI
if [[ "$method" == "all" ]]; then
  echo Compiling and executing golang build...
  go build && go run main.go
fi
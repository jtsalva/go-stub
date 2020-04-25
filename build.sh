declare -a builds=("linux amd64" "linux arm64" "darwin amd64" "windows amd64")

rm -rf ./bin
for build in "${builds[@]}"
do
  read -r -a parts <<< "$build"
  os="${parts[0]}"
  arch="${parts[1]}"
  echo "Building $os/$arch"
  GOOS=$os GOARCH=$arch go build -o "bin/$os/$arch/go-stub" config.go middleware.go stub.go yaml_structs.go main.go
done

echo "Finished building into ./bin"
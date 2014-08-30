build-darwin:
	GOARCH=amd64 GOOS=darwin go build -v -o pkg/darwin/subpro

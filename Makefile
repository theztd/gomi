build-linux:
	GOOS=linux GOARCH=amd64 go build -o build/linux/gomi .
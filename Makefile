
build:
	go build -mod vendor -o sdk

release:
	GOOS=linux CGO_ENABLED=0 go build -mod vendor -o sdk

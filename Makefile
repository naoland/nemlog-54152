
version:
	@go version
init:
	go mod init notif
run:
	go run notif.go -token ${LINE_TOKEN}  テストメッセージです
build:
	go build -o dist/notif notif.go
test:
	go test -v -cover -timeout 30s notif_test.go -run ^TestHello$
build-windows:
	echo "Build for windows10"
	GOOS=windows GOARCH=amd64 go build -o dist/windows/notif.exe notif.go
	echo "Done!"
build-linux:
	echo "Build for linux"
	GOOS=linux GOARCH=amd64 go build -o dist/linux/notif notif.go
	echo "Done!"
build-mac:
	echo "Build for macOS(Darwin)"
	GOOS=darwin GOARCH=amd64 go build -o dist/macos/notif notif.go
	echo "Done!"
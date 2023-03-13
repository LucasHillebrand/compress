install:
	go build -o comp .
	mv comp /usr/local/bin

build:
	mkdir bin
	GOOS=windows GOARCH=amd64 go build -o bin/comp\ Arch=amd64\ OS=windows.exe .
	GOOS=windows GOARCH=386 go build -o bin/comp\ Arch=386\ OS=windows.exe .
	GOOS=darwin GOARCH=amd64 go build -o bin/comp\ Arch=amd64\ OS=darwin .
	GOOS=darwin GOARCH=arm64 go build -o bin/comp\ Arch=arm64\ OS=darwin .
	GOOS=linux GOARCH=amd64 go build -o bin/comp\ Arch=amd64\ OS=linux .
	GOOS=linux GOARCH=386 go build -o bin/comp\ Arch=386\ OS=linux .
	GOOS=linux GOARCH=arm64 go build -o bin/comp\ Arch=arm64\ OS=linux .
	GOOS=linux GOARCH=arm go build -o bin/comp\ Arch=arm\ OS=linux .
build:
	GOOS=windows GOARCH=amd64 go build -o Text_Adventure_Windows.exe
	GOOS=linux GOARCH=amd64 go build -o Text_Adventure_Linux
	GOOS=darwin GOARCH=amd64 go build -o Text_Adventure_Mac

deploy: build
	mv ./Text_Adventure_Windows.exe ./bin/Windows/Text_Adventure.exe
	mv ./Text_Adventure_Linux ./bin/linux/Text_Adventure
	mv ./Text_Adventure_Mac ./bin/mac/Text_Adventure

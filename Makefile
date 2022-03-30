build:
	GOOS=windows GOARCH=amd64 go build -o Text_Adventure_Windows.exe
	GOOS=linux GOARCH=amd64 go build -o Text_Adventure_Linux
	GOOS=darwin GOARCH=amd64 go build -o Text_Adventure_Mac
	mv ./Text_Adventure_Windows.exe ./bin/Windows.exe
	mv ./Text_Adventure_Linux ./bin/linux
	mv ./Text_Adventure_Mac ./bin/mac

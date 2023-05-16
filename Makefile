default:
	go run .

build:
	go build .

release:
	GOOS=linux GOARCH=amd64 go build .
	mkdir boarder_linux64/
	mv boarder boarder_linux64/
	cp template.html boarder_linux64/
	touch boarder_linux64/threads.txt
	zip -r boarder_linux64.zip boarder_linux64/
	sha256sum boarder_linux64.zip > boarder_linux64.sha256sum
	rm -rf boarder_linux64/
	GOOS=windows GOARCH=amd64 go build .
	mkdir boarder_win64
	mv boarder.exe boarder_win64/
	cp template.html boarder_win64/
	touch boarder_win64/threads.txt
	zip -r boarder_win64.zip boarder_win64/
	sha256sum boarder_win64.zip > boarder_win64.sha256sum
	rm -rf boarder_win64/

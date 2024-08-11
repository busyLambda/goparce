test:
	yes hello1 | head -n 1000000 > output.txt
	go test ./...


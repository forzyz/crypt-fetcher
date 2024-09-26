build:
	go build -o bin/cryptfetcher 

run: build 
	./bin/cryptfetcher
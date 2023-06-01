install:
	go build -o /tmp/chains

run: install
	/tmp/chains

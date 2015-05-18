install:
	go build -o /usr/local/bin/ghcnt github_count.go

uninstall:
	rm /usr/local/bin/ghcnt

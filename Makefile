clean:
	rm -rf dist/.tmpcntr dist/cgi-bin/ulticntr

build:
	@mkdir -p dist/cgi-bin
	go build -o dist/cgi-bin/ulticntr ulticntr/*.go
	go build -o dist/devserver


run: clean build
	./dist/devserver

.PHONY: dev
dev:
	ls ./*.go ulticntr/*.go ulticntr/assets/*.go.html | entr -r make run

docker-test: test_mode=true
docker-run: test_mode=false
docker-run docker-test:
	GOARCH=amd64 GOOS=linux make build && \
	docker-compose build && \
	TEST_MODE=$(test_mode) docker-compose up \
	--exit-code-from cgibinftw
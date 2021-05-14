build:
	@mkdir -p dist/cgi-bin
	go build -o dist/cgi-bin/ulticntr ulticntr/*.go

run: build
	go run main.go

.PHONY: dev


dev dev-tmp:
	ls ./*.go ulticntr/*.go ulticntr/assets/*.go.html | entr -r make run

docker-test: test_mode=true
docker-run: test_mode=false
docker-run docker-test:
	docker-compose build && \
	TEST_MODE=$(test_mode) docker-compose up \
	--exit-code-from cgibinftw
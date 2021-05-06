
build:
	go build -o dist/cgi-bin/ulticntr ulticntr/*.go

run: build
	go run main.go

# dev uses enter in reload mode to watch changes in go files and template files
# to rebuild ulticntr and restart the cgi server
dev:
	ls ./*.go ulticntr/*.go ulticntr/assets/*.go.html | entr -r make run
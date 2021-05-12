FROM golang:1.16.4-alpine3.13

WORKDIR /cgibinftw
EXPOSE 8888

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir -p dist/cgi-bin && \
    go build -o dist/cgi-bin/ulticntr ulticntr/*.go && \
    go build

CMD [ "./entrypoint.sh" ]

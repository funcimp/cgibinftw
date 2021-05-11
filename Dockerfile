FROM golang:latest

WORKDIR /cgibinftw

COPY . .

EXPOSE 8888

CMD [ "make", "run" ]

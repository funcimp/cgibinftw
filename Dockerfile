FROM scratch

WORKDIR /cgibinftw
EXPOSE 8888

COPY dist dist

CMD [ "./dist/devserver" ]

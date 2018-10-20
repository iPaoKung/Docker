FROM golang:latest
WORKDIR /go/src/hello
COPY ./src /go/src/hello
RUN ls -l /go/src/hello
RUN go install /go/src/hello
CMD ["hello"]
EXPOSE 8080

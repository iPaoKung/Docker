FROM golang:latest
WORKDIR /go/src/paolog
COPY ./src /go/src/paolog
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/labstack/echo
RUN ls -l /go/src/paolog
RUN go install /go/src/paolog
CMD ["paolog"]
EXPOSE 8080

FROM golang:1.10

WORKDIR $GOPATH/src

COPY . .

RUN go get -d -v ./...
RUN go build -o ./out/app1 .

EXPOSE 8080
USER 0

CMD ["./out/app1"]

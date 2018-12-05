# FROM arm32v6/golang:alpine
FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
# RUN go install -v ./...
RUN env GOOS=linux GOARCH=arm GOARM=6 go build

CMD ["./app"]

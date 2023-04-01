FROM golang:1.16-alpine
MAINTAINER Vikas Sharma
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /main.go

EXPOSE 8080

CMD [ ".main.exe" ]
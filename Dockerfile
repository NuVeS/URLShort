FROM golang:1.18 AS build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go build -o main . 
CMD ["/app/main"]

ENV CGO_ENABLED=0
RUN go get -d -v ./...

FROM scratch AS runtime
ADD URLShort /
CMD ["/main"]
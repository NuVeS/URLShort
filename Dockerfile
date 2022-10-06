FROM golang:1.18 AS build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY cmd ./

RUN go build -o main . 
CMD ["/URLShort"]
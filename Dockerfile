FROM golang:1.21.0

WORKDIR /wbsapp

COPY go.mod ./
COPY go.sum ./
RUN go get -t -v ./...
RUN go mod download

COPY *.go ./

RUN mkdir ./static
COPY static ./static

RUN mkdir ./pages
COPY pages ./pages

COPY english_cleaned.txt ./


RUN go build -o ./wbsapp

EXPOSE 80

CMD ["/wbsapp/wbsapp"]
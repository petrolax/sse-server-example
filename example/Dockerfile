FROM golang:latest

WORKDIR server

COPY . .

RUN go build -o server main.go

CMD ["./server"]
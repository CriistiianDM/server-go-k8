#Golang image
FROM golang:1.23.0

WORKDIR /server_go

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o server ./main.go

EXPOSE 4000

CMD ["./server"]
FROM golang:latest

WORKDIR /app

COPY . .

RUN go get github.com/elastic/go-sysinfo/providers/linux@v1.1.1

RUN  go build

CMD ["./Parking-lot-Service"]

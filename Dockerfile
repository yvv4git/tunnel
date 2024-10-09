FROM golang:1.23

RUN apt update && apt install -y iproute2 net-tools netcat-openbsd vim

WORKDIR /app

COPY . .

RUN go build -o tunnel main.go

CMD ["./run.sh"]
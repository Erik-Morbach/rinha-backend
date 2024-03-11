FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download

#RUN go test

RUN go build -o /rinha

EXPOSE 8080

CMD ["sh", "-c", "/rinha"]

FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main cmd/*

EXPOSE 1234

CMD [ "./main" ]
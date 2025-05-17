FROM --platform=linux/amd64 golang:1.24

WORKDIR /go/src/app

COPY . .

EXPOSE 8001

RUN go build -o main cmd/main.go

CMD [ "./main" ]
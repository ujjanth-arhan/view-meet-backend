FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./view-meet-backend

EXPOSE 8080

CMD [ "./view-meet-backend" ]
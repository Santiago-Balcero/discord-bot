FROM golang:1.22.0

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/src/app ./...

CMD ["/usr/src/app/discord-spotify-bot"]
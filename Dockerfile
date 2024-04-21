FROM golang:1.22.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN git config --global --add safe.directory /usr/src/app
RUN go build -v -o /app ./...

CMD ["/app/discord-spotify-bot"]

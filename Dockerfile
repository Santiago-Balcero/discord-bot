FROM golang:1.22.0

RUN useradd -m -s /bin/bash appuser
USER appuser

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN git config --global --add safe.directory /usr/src/app
RUN go build -v -o /usr/src/app ./...

RUN rm -rf go.mod go.sum *.go .env .gitignore .todo *.md

CMD ["./bin/discord-spotify-bot"]

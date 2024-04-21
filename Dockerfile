FROM golang:1.22.0 as build

COPY . /app
WORKDIR /app

# Setup buildpack
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://buildpack-registry.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

#Execute Buildpack
RUN STACK=heroku-20 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

FROM golang:1.22.0

COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m botuser
USER botuser

CMD /app/bin/discord-spotify-go

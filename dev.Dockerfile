FROM golang:1.18.4-alpine

RUN apk update && \
  apk add make git bash curl openssl alpine-sdk --no-cache

WORKDIR /app

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
  && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

CMD air

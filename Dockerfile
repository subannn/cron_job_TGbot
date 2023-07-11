FROM golang

WORKDIR /app

COPY . .

ENV BOT_KEY=       # bot key
ENV SUPER_USER_ID= # id

RUN go build main.go

ENTRYPOINT /app/main



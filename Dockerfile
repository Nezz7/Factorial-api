FROM golang

WORKDIR /app

COPY . .

RUN go build -o /app/service fact.go


CMD '/app/service'
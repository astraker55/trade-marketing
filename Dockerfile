FROM golang:latest

LABEL maintainer = "Hasratian Konstantin astraker98@gmail.com"

ADD dbdump.sql /docker-entrypoint-initdb.d/

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build .

EXPOSE 9200

ENTRYPOINT [ "./trade-marketing" ]
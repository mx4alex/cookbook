FROM golang:1.20

WORKDIR /usr/src/app

RUN apt-get update \
    && apt-get -y install postgresql-client

COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

RUN chmod +x entrypoint.sh \
    && go build -o taskmanager ./cmd/app/main.go

CMD ["./cookbook"]
FROM golang:latest

RUN mkdir src/app
WORKDIR /src/app
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build ./cmd/web/
CMD ["./web"]
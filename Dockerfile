FROM golang:latest

RUN go version
ENV GOPATH=/


COPY ./ ./

RUN apt-get update && apt-get -y install postgresql-client
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN chmod +x wait-for-postgres.sh


RUN go mod download
RUN go build -o main main.go


EXPOSE 8080


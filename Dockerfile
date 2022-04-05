FROM golang:1.18-alpine

WORKDIR /app
COPY . ./

RUN go mod download
RUN go build -o /user-balance ./cmd

EXPOSE 8080

ENTRYPOINT ["/user-balance"]
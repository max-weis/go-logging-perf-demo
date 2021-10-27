FROM golang:1.17.1-alpine as build

WORKDIR $GOPATH/app/

RUN apk add git

COPY go.* ./
RUN go mod download

COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o logging-perf-log

#resulting app
FROM alpine as final
COPY --from=build go/app/logging-perf-log /app/
WORKDIR /app
RUN mkdir log

ENTRYPOINT [ "./logging-perf-log" ]
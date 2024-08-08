FROM golang:alpine3.18 as builder
RUN apk add --no-cache --update git build-base

WORKDIR /app
COPY . .
RUN go mod tidy && go build \
    -a \
    -trimpath \
    -o annona_bot \
    -ldflags "-s -w -buildid=" \
    "./cmd/annona_bot" && \
    ls -lah

FROM alpine:3.18 as runner
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/annona_bot .
VOLUME /app/log
#EXPOSE 8080

ENTRYPOINT ["./annona_bot"]
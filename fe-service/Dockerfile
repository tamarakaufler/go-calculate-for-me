FROM alpine:3.7
RUN apk add --no-cache openssh ca-certificates

RUN mkdir /app
WORKDIR /app
COPY fe-service /app
ENTRYPOINT ["./fe-service"]
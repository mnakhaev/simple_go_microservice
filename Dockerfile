FROM golang:1.19-alpine3.16 AS build-env

# Allow Go to retrieve the dependencies for the build step
RUN apk add --no-cache git

# Secure against running as root
RUN adduser -D -u 10000 mn
RUN mkdir /microservice/ && chown mn /microservice/
USER mn

WORKDIR /microservice/
ADD . /microservice/

# Compile the binary, we don't want to run the cgo resolver
RUN CGO_ENABLED=0 go build -o  /microservice/mcrsrv .

# Final stage
FROM alpine:3.8

# Secure against running as root
RUN adduser -D -u 10000 mn
USER mn

WORKDIR /
COPY --from=build-env /microservice/mcrsrv /

EXPOSE 8080

CMD ["/mcrsrv"]

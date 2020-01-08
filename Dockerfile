FROM golang

# Fetch dependencies
RUN go get github.com/tools/godep

# Add project directory to Docker image.
ADD . /go/src/github.com/geordie/toplin

ENV USER geordie
ENV HTTP_ADDR :8888
ENV HTTP_DRAIN_INTERVAL 1s
ENV COOKIE_SECRET FWnQhq62ylAMfOAH

# Replace this with actual PostgreSQL DSN.
ENV DSN postgres://geordie@localhost:5432/toplin?sslmode=disable

WORKDIR /go/src/github.com/geordie/toplin

RUN godep go build

EXPOSE 8888
CMD ./toplin
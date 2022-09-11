FROM golang:latest as BUILD
WORKDIR tests
COPY . .
ENTRYPOINT ["go", "test", "-v", "./..."]
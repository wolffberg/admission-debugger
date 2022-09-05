#FROM golang:1.19-bullseye as build
FROM golang:alpine as build

WORKDIR /src

COPY ./ca-certificates.crt /etc/ssl/certs/

COPY ./admission-debugger /src

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

RUN go test ./...

RUN go build -ldflags="-w -s" .

#FROM scratch
FROM golang:1.19-bullseye

COPY --from=build /src/admission-debugger /src/

WORKDIR /src

ENTRYPOINT ["/src/admission-debugger"]

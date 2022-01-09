FROM golang:alpine AS build
ENV GOPROXY=https://proxy.golang.org
RUN apk add --update git
WORKDIR /go/src/api
COPY . .
RUN GOOS=linux go build -o /go/bin/api cmd/api/main.go

FROM alpine
COPY --from=build /go/bin/api /go/bin/api
ENTRYPOINT ["go/bin/api"]
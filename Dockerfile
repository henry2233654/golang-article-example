FROM golang:latest AS build-env
WORKDIR $GOPATH/src/golang-article-example
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -v -installsuffix cgo .

FROM scratch
ENV IS_CONTAINER=TRUE
COPY --from=build-env go/src/golang-article-example/golang-article-example .
EXPOSE 8081
CMD ["./golang-article-example"]
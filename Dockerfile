FROM golang:1.16-alpine AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /workspace
COPY . .
RUN go mod download -x && go build -a -o /main .

FROM golang:1.16-alpine

WORKDIR /
RUN mkdir /.resources
COPY --from=build /main /main
#COPY src/app/.resources/*.properties /.resources
EXPOSE 8080
CMD ["/main", "serve"]

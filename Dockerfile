FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /ttdhbw

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /ttdhbw /ttdhbw

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/ttdhbw"]

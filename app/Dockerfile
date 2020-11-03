FROM golang AS build

WORKDIR /app

COPY . .

ENV \
  GOPATH=/app \
  CGO_ENABLED=0 \
  GOOS=linux

RUN go test ./...

RUN go install -a -ldflags '-extldflags "-static"' ./...

FROM scratch

WORKDIR /app

COPY --from=build /app .

ENTRYPOINT [ "./bin/frutimo" ]

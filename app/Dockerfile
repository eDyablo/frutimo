FROM golang AS build

WORKDIR /app

COPY . .

ENV \
  GOPATH=/app \
  CGO_ENABLED=0 \
  GOOS=linux

RUN cat src/packages | tr $'\n' ' ' | xargs go get -d

RUN go test ./...

RUN go install -a -ldflags '-extldflags "-static"' ./...

FROM scratch

WORKDIR /app

COPY --from=build /app .

ENTRYPOINT [ "./bin/frutimo" ]

EXPOSE 80

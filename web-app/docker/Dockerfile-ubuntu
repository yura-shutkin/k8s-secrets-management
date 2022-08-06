ARG GOLANG_VER=1.18.2
FROM golang:${GOLANG_VER}-alpine as build
WORKDIR /app
ENV CGO_ENABLED=0
COPY src/. .
RUN go mod tidy
RUN go build -o web-app && chmod a+x web-app

FROM ubuntu:20.04
RUN adduser --no-create-home --disabled-password web-app
USER web-app
WORKDIR /app
COPY --from=build /app/web-app /app/web-app
COPY --from=build /app/*.gohtml /app/
ENTRYPOINT ["/app/web-app"]

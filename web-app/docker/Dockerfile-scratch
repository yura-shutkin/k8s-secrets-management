ARG GOLANG_VER=1.18.2
FROM golang:${GOLANG_VER}-alpine as build
WORKDIR /app
ENV CGO_ENABLED=0
COPY src/. .
RUN go mod tidy
RUN go build -o web-app && chmod a+x web-app
RUN adduser -HD web-app

FROM scratch
COPY --from=build /etc/passwd /etc/passwd
USER web-app
WORKDIR /app
COPY --from=build /app/web-app /app/web-app
COPY --from=build /app/*.gohtml /app/
ENTRYPOINT ["/app/web-app"]

FROM golang:1.22 AS build
WORKDIR /src
COPY . .
RUN go build -o /out/tile-server ./cmd/server
FROM gcr.io/distroless/base-debian12
COPY --from=build /out/tile-server /tile-server
EXPOSE 18114
ENTRYPOINT ["/tile-server"]

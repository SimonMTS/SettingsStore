FROM golang:1.19 AS build

WORKDIR /app
RUN wget -q https://github.com/go-swagger/go-swagger/releases/download/v0.30.3/swagger_linux_amd64
RUN chmod +x ./swagger_linux_amd64
RUN mv ./swagger_linux_amd64 /bin/swagger
COPY . .
RUN make generate
RUN go build -o setting-store .

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /app/setting-store /setting-store
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/setting-store"]

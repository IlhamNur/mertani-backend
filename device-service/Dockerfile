FROM golang:1.25-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/app .
EXPOSE 8080
CMD ["./app"]
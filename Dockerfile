# Build Stage
FROM golang:1.21.1-alpine3.18 As BuildStage

WORKDIR /app
COPY . /app
EXPOSE 8080
RUN go build -o /main ./console/main.go


## Deploy Stage
FROM alpine
WORKDIR /
COPY --from=BuildStage /main /main
COPY ./config/files /app/config/files
EXPOSE 8080
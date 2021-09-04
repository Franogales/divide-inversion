FROM golang:1.17-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY second/ ./
COPY second/service /
RUN ls -la
RUN pwd
EXPOSE 8080
CMD ["/service"]
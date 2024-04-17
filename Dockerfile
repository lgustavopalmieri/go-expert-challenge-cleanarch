FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY .env ./ 
RUN apk add --no-cache make 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem/main.go


FROM alpine:latest
RUN apk --no-cache add ca-certificates make
WORKDIR /app
COPY --from=builder /app/main /main
COPY --from=builder /app/.env ./  
COPY --from=builder /app/cmd/ordersystem/sql/migrations /app/cmd/ordersystem/sql/migrations
EXPOSE 8086

ENTRYPOINT ["/main"]




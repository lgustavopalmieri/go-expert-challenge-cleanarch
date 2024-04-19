# go-expert-challenge-cleanarch

Este é o código do desafio do módulo Go Expert: Clean Architecture.

## rest

O serviço rest roda na porta 8086.
Para utilizá-lo você pode acessar na pasta "api" os dois endpoints, ou se optar por outra ferramenta:

```http
POST http://localhost:8086/orders/create
Content-Type: application/json
{
    "price": 333,
    "tax": 33.0
}
```

```http
GET http://localhost:8086/orders/list
Content-Type: application/json
```

## grpc

Para utilizar o serviço grpc é preciso ter o evans instalado.

```
go install github.com/ktr0731/evans@latest
```

Depois de instalado você pode rodar os comandos na sequência:

```
evans -r repl
package orderpb
service OrderService
call CreateOrder
call ListOrders
```

## graphql

Para utilizar o serviço graphql basta entrar em http://localhost:8282/ e colar:

```graphql
mutation createOrder {
  createOrder(input: { Price: 50.4, Tax: 3.0 }) {
    OrderID
    Price
    Tax
    FinalPrice
    CreatedAt
  }
}

query queryOrders {
  orders {
    OrderID
    Price
    Tax
    FinalPrice
    CreatedAt
  }
}
```

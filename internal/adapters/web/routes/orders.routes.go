package routes

import (
	"database/sql"

	orders "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/handlers/order"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/server"
	_ "github.com/lib/pq"
)

func SetupOrdersRoutes(s *server.WebServer, db *sql.DB) {
	webOrderHandler := orders.NewWebOrderHandler(db)
	s.AddHandler("/orders/create", webOrderHandler.Create)
	s.AddHandler("/orders/list", webOrderHandler.ListOrders)
}

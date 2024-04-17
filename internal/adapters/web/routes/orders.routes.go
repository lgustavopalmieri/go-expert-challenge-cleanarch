package routes

import (
	"database/sql"

	orders "github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/handlers/order"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/internal/adapters/web/server"
	"github.com/lgustavopalmieri/go-expert-challenge-cleanarch/pkg/events"
	_ "github.com/lib/pq"
)

func SetupOrdersRoutes(s *server.WebServer, db *sql.DB, eventDispatcher events.EventDispatcherInterface) {
	webOrderHandler := orders.NewWebOrderHandler(db, eventDispatcher)
	s.AddHandler("/orders/create", webOrderHandler.Create)
	s.AddHandler("/orders/list", webOrderHandler.ListOrders)
}

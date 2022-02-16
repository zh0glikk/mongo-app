package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/zh0glikk/mongo-app/internal/api/handlers"
	"github.com/zh0glikk/mongo-app/internal/api/middlewares"
	mongo2 "github.com/zh0glikk/mongo-app/internal/data/mongo"
)

func Router(log *logrus.Entry, mongoCli *mongo.Client) (chi.Router, error) {
	r := chi.NewRouter()

	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middlewares.CtxMiddlewares(
			handlers.CtxLog(log),
			handlers.CtxItems(mongo2.NewItems(mongoCli)),
		),
	)

	r.Route("/items", func(r chi.Router) {
		r.Post("/", handlers.AddItem)
		r.Get("/", handlers.GetItems)

		r.Route("/{id}", func(r chi.Router) {
			r.Put("/", handlers.UpdateItem)
			r.Delete("/", handlers.DeleteItem)
			r.Get("/", handlers.GetItem)
		})
	})

	return r, nil
}

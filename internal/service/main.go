package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/zh0glikk/mongo-app/internal/api"
	"github.com/zh0glikk/mongo-app/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	cfg *config.Config
	log *logrus.Entry
}

func NewService(cfg *config.Config) *Service {
	switch {
	case cfg == nil:
		panic("cfg is nil")
	}

	return &Service{
		cfg: cfg,
		log: logrus.NewEntry(logrus.New()),
	}
}

func (s *Service) Run(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.cfg.MongoCfg.URL))
	if err != nil {
		return errors.Wrap(err, "failed to init mongo client")
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	router, err := api.Router(s.log, client)
	if err != nil {
		return err
	}

	return http.ListenAndServe(s.cfg.ListenerConfig.Addr, router)
}

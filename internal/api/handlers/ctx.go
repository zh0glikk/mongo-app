package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/zh0glikk/mongo-app/internal/data"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	itemsQCtxKey
)

func CtxLog(entry *logrus.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logrus.Entry {
	return r.Context().Value(logCtxKey).(*logrus.Entry)
}

func CtxItems(entry data.Items) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, itemsQCtxKey, entry)
	}
}

func Items(r *http.Request) data.Items {
	return r.Context().Value(itemsQCtxKey).(data.Items)
}
//
//func CtxMongo(entry *mongo.Client) func(context.Context) context.Context {
//	return func(ctx context.Context) context.Context {
//		return context.WithValue(ctx, mongoCtxKey, entry)
//	}
//}
//
//func Mongo(r *http.Request) *mongo.Client {
//	return r.Context().Value(mongoCtxKey).(*mongo.Client)
//}
//
//

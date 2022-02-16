package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
//
type Items interface {
	Create(ctx context.Context, item Item) (string, error)
	Delete(ctx context.Context, id string) (*Item, error)
	Get(ctx context.Context, id string) (*Item, error)
	Update(ctx context.Context, id string, item Item) error

	FilterByAmountGt(amount uint64) Items
	FilterByAmountLt(amount uint64) Items
	FilterByTitle(title string) Items


	ClearFilters() Items
	GetBatch(ctx context.Context) ([]Item, error)
}

type Item struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title  string             `json:"name"`
	Amount uint64             `json:"amount"`
	Price  uint64             `json:"price"`
}

package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/zh0glikk/mongo-app/internal/data"
)

var itemsCollection = "items"

type Items struct {
	filters    map[string]interface{}
	collection *mongo.Collection
}

func NewItems(mongoCli *mongo.Client) data.Items {
	return &Items{
		filters:    make(map[string]interface{}),
		collection: mongoCli.Database(mongoDB).Collection(itemsCollection),
	}
}

func (i *Items) ClearFilters() data.Items {
	i.filters = map[string]interface{}{}

	return i
}

func (i *Items) Create(ctx context.Context, item data.Item) (string, error) {
	res, err := i.collection.InsertOne(ctx, &item)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (i *Items) Delete(ctx context.Context, id string) (*data.Item, error) {
	var item data.Item

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = i.collection.FindOneAndDelete(ctx, bson.M{"_id": objectId}).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (i *Items) Get(ctx context.Context, id string) (*data.Item, error) {
	var item data.Item

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = i.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&item)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (i *Items) Update(ctx context.Context, id string, item data.Item) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = i.collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.D{
		{"$set", item},
	})

	return err
}

func (i *Items) FilterByAmountGt(amount uint64) data.Items {
	_, ok := i.filters["amount"]
	if ok {
		i.filters["amount"].(bson.M)["$gt"] = amount

		return i
	}

	i.filters["amount"] = bson.M{"$gt": amount}

	return i
}

func (i *Items) FilterByAmountLt(amount uint64) data.Items {
	_, ok := i.filters["amount"]
	if ok {
		i.filters["amount"].(bson.M)["$lte"] = amount

		return i
	}

	i.filters["amount"] = bson.M{"$lte": amount}

	return i
}

func (i *Items) FilterByTitle(title string) data.Items {
	_, ok := i.filters["title"]
	if ok {
		i.filters["title"] = title

		return i
	}

	i.filters = bson.M{"title": title}

	return i
}

func (i *Items) GetBatch(ctx context.Context) ([]data.Item, error) {
	items := make([]data.Item, 0)

	cur, err := i.collection.Find(ctx, i.filters)
	if err != nil {
		return items, err
	}

	err = cur.All(ctx, &items)
	if err != nil {
		return items, err
	}

	return items, nil
}

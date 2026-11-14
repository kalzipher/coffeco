package purchase

import (
	"context"
	"time"

	coffeco "github.com/kalzipher/coffeco/internal"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository interface {
	Store(ctx context.Context, purchase *Purchase) error
}

/* Mongo implementations */
type MongoRepository struct {
	purchase *mongo.Collection
}

func NewMongoRepository(connectionString string) (*MongoRepository, error) {

	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	purchase := client.Database("purchase").Collection("purchases")

	// Implementation for creating MongoDB connection and collection
	return &MongoRepository{
		purchase: purchase,
	}, nil
}

func (r *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoP := new(purchase)
	_, err := r.purchase.InsertOne(ctx, mongoP)
	if err != nil {
		return err
	}

	return nil
}

type mongoPurchase struct {
	ID                 string            `bson:"id"`
	Store              string            `bson:"store"`
	productsToPurchase []coffeco.Product `bson:"products_to_purchase"`
	Total              int64             `bson:"total"`
	timeOfPurchase     string            `bson:"time_of_purchase"`
	cardToken          *string           `bson:"card_token,omitempty"`
}

func toMongoPurchase(purchase *Purchase) *mongoPurchase {
	return &mongoPurchase{
		ID:                 purchase.id.String(),
		Store:              purchase.Store.ID.String(),
		productsToPurchase: purchase.ProductToPurchase,
		Total:              purchase.total.Amount(),
		timeOfPurchase:     purchase.timeOfPurchase.Format(time.RFC3339),
		cardToken:          purchase.CardToken,
	}
}

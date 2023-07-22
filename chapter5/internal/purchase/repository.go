package purchase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	coffeeco "github.com/jhseoeo/Golang-DDD/chapter5/internal"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/payment"
	"github.com/jhseoeo/Golang-DDD/chapter5/internal/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Repository interface {
	Store(ctx context.Context, purchase Purchase) error
}

type MongoRepository struct {
	purchases *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purchases := client.Database("coffeeco").Collection("purchases")

	return &MongoRepository{
		purchases: purchases,
	}, nil
}

func (mr *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoP := toMongoPurchase(purchase)
	_, err := mr.purchases.InsertOne(ctx, mongoP)
	if err != nil {
		return fmt.Errorf("failed to persist purchase: %w", err)
	}

	return nil
}

type mongoPurchase struct {
	id                 uuid.UUID          `bson:"ID"`
	store              store.Store        `bson:"Store"`
	productsToPurchase []coffeeco.Product `bson:"product_purchased"`
	total              int64              `bson:"purchase_total"`
	paymentMeans       payment.Means      `bson:"payment_means"`
	timeOfPurchase     time.Time          `bson:"created_at"`
	cardToken          *string            `bson:"card_token"`
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		id:                 p.id,
		store:              p.Store,
		productsToPurchase: p.ProductsToPurchase,
		total:              int64(p.total),
		paymentMeans:       p.PaymentMeans,
		timeOfPurchase:     p.timeOfPurchase,
		cardToken:          p.CardToken,
	}
}

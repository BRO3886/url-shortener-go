package database

import (
	"context"
	"time"

	"github.com/BRO3886/url-shortener/shortener"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepo struct {
	db      *mongo.Database
	dbName  string
	timeout time.Duration
}

// Create implements shortener.RedirectRepository
func (r *mongoRepo) Create(redirect *shortener.Redirect) (*shortener.Redirect, error) {
	collection := r.db.Collection("redirects")
	_, err := collection.InsertOne(context.Background(), redirect)
	if err != nil {
		return nil, err
	}

	return redirect, nil
}

// Find implements shortener.RedirectRepository
func (r *mongoRepo) Find(code string) (*shortener.Redirect, error) {
	redirect := new(shortener.Redirect)
	collection := r.db.Collection("redirects")
	filter := bson.M{"code": code}

	err := collection.FindOne(context.Background(), filter).Decode(redirect)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, shortener.ErrRedirectNotFound
		}
		return nil, err
	}

	return redirect, nil
}

func NewRepository(mongoURL, database string, timeout int) (shortener.RedirectRepository, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Duration(timeout)*time.Second)
	// defer cancel()
	ctx := context.Background()

	client, err := newMongoClient(ctx, mongoURL)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database(database)

	return &mongoRepo{
		db:      db,
		dbName:  database,
		timeout: time.Duration(timeout) * time.Second,
	}, nil
}

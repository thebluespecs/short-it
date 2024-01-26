package mongo

import (
	"context"
	"errors"
	"short-it/config"
	"short-it/internal/db/mongo/models"
	"short-it/internal/logger"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo is a struct that implements DB interface
type Mongo struct {
    Client *mongo.Client
    Collection *mongo.Collection
}

var (
	instance *Mongo
	one      sync.Once
)

func init() {
	one.Do(func() {
        instance = &Mongo{
            Client: nil,
            Collection: nil,
        }
        instance.Connect()
	})
}

func GetInstance() *Mongo {
	return instance
}

func (m *Mongo) Connect() error {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Get("MONGO_URL")))
	if err != nil {
		panic(err)
	}
    defer func() {
        if err = client.Disconnect(context.Background()); err != nil {
            panic(err)
        }
    }()
    m.Client = client
    m.Collection = client.Database("short-it").Collection("urls")
    return nil
}

func (m *Mongo) Disconnect() error {
    if m.Client == nil {
        return errors.New("client not initialized")
    }
    return m.Client.Disconnect(context.Background())
}

// Save function to save the url to the mongo database to the collection urls
func (m *Mongo) Save(url string, expiresAt time.Duration) (int, error) {
    shortUrl := models.NewShortUrl(url, expiresAt)
    result, err := m.Collection.InsertOne(context.TODO(), shortUrl)
    if err != nil {
        logger.Error(err.Error())
        return 0, err
    }
    logger.Info("inserted id: " + result.InsertedID.(string))
    return int(result.InsertedID.(int64)), nil
}

// Find finds the url from the database
func (m *Mongo) Find(id int) (string, error) {
    var url models.ShortUrl
    err := m.Collection.FindOne(context.TODO(), bson.M{
        "_id": id,
    }).Decode(&url)
    if err != nil {
        logger.Error(err.Error())
        return "", err
    }
    return url.Url, nil
}

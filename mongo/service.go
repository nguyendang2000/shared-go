package mongo

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Service struct contains the MongoDB client and context
type Service struct {
	client *mongo.Client
	ctx    context.Context
}

// NewService initializes a new MongoDB connection using the given context and configuration
func NewService(ctx context.Context, conf Config) (*Service, error) {
	fullAddress := conf.Address

	if !strings.HasPrefix(fullAddress, "mongodb://") {
		fullAddress = "mongodb://" + fullAddress
	}

	clientOptions := options.Client().ApplyURI(fullAddress)
	if conf.Username != "" && conf.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username:   conf.Username,
			Password:   conf.Password,
			AuthSource: conf.AuthDB,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToConnect, err)
	}

	// Use a separate context for the ping operation to avoid shadowing
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := client.Ping(pingCtx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf(ErrFailedToPing, err)
	}

	return &Service{
		client: client,
		ctx:    ctx,
	}, nil
}

// Client returns the MongoDB client from the Service struct
func (inst *Service) Client() *mongo.Client {
	return inst.client
}

// Database returns a MongoDB database instance by name
func (inst *Service) Database(dbName string) *mongo.Database {
	return inst.client.Database(dbName)
}

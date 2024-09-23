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

// Service struct contains the MongoDB client and a timeout field
type Service struct {
	client  *mongo.Client
	timeout int64 // Timeout in seconds for requests
}

// NewService initializes a new MongoDB connection using the given configuration
// and sets the timeout in the Service struct.
func NewService(conf Config) (*Service, error) {
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

	// Set timeout to DefaultTimeout if not provided or less than 0
	timeout := conf.Timeout
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	// Create a context with the specified timeout
	timeoutDuration := time.Duration(timeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToConnect, err)
	}

	// Ping the primary MongoDB node to verify connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf(ErrFailedToPing, err)
	}

	return &Service{client: client, timeout: timeout}, nil
}

// Close closes the MongoDB client connection
func (inst *Service) Close(ctx context.Context) error {
	if err := inst.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to close MongoDB connection: %v", err)
	}
	return nil
}

// Ping checks if MongoDB is still available
func (inst *Service) Ping(ctx context.Context) error {
	if err := inst.client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf(ErrFailedToPing, err)
	}
	return nil
}

// Database returns a MongoDB database instance by name
func (inst *Service) Database(dbName string) *mongo.Database {
	return inst.client.Database(dbName)
}

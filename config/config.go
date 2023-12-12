package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

type Config struct {
	MongoOptions MongoOptions `yaml:"mongo"`
}

type MongoOptions struct {
	ConnectionString string `yaml:"connectionString"`
	Database         string `yaml:"database"`
}

func GetFromYaml() (*Config, error) {
	config := &Config{}

	file, err := os.Open("../config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func (config Config) GetMongoDb(ctx context.Context) (*mongo.Database, error) {
	mongoOptions := options.Client().ApplyURI(config.MongoOptions.ConnectionString)
	mongoOptions.TLSConfig.InsecureSkipVerify = true

	mongoClient, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return mongoClient.Database(config.MongoOptions.Database), nil
}

package repositories

import (
	"context"

	"github.com/LastSprint/InfrastructureSlackApp/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBContext контекст для работы с БД.
type DBContext struct {
	client *mongo.Client
	db     *mongo.Database
	cntx   context.Context
}

// NewDB создает новый контекст и подключается к базе
func NewDB() (*DBContext, error) {
	return createDBContext(utils.Config.MongoDBConfig.ConnectionString, utils.Config.MongoDBConfig.DataBaseString)
}

// NewTestDB создает контекст для тестов и подключается к тестовой базе.
func NewTestDB() (*DBContext, error) {
	return createDBContext(utils.Config.MongoDBConfig.TestConnectionString, utils.Config.MongoDBConfig.TestDataBaseString)
}

// Close закрывает подключение к базе данных.
func (repo *DBContext) Close() error {
	return repo.client.Disconnect(repo.cntx)
}

func createDBContext(cnstr string, dbstr string) (*DBContext, error) {
	connectionClient := options.Client().ApplyURI(cnstr)
	client, err := mongo.NewClient(connectionClient)
	cntx := context.TODO()

	if err != nil {
		return nil, err
	}

	err = client.Connect(cntx)

	if err != nil {
		return nil, err
	}

	db := client.Database(dbstr)

	return &DBContext{client: client, db: db, cntx: cntx}, nil
}

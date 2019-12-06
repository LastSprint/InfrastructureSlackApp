package repositories

import (
	"context"

	"github.com/LastSprint/InfrastructureSlackApp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "users"

// UserRepository репозиторий для работы с пользователями `User`.
type UserRepository interface {
	// ReadAllDevelopers Читает всех разработчиков - member.lead, member.developer
	ReadAllDevelopers() ([]*models.User, error)
	// ReadDevelopers Читает только разработчиков - member.developer
	ReadDevelopers() ([]*models.User, error)
	// ReadLeadDevelopers считывает всех лид-разработчиков - member.lead
	ReadLeadDevelopers() ([]*models.User, error)
	// ReadManagers считывает всех менеджеров - member.manager
	ReadManagers() ([]*models.User, error)
}

// UserDBRepository репозиторий с доступом к БД.
type UserDBRepository struct {
	DB *DBContext
}

// ReadAllDevelopers Читает всех разработчиков - member.lead, member.developer
func (rep *UserDBRepository) ReadAllDevelopers() ([]*models.User, error) {

	filter := bson.M{
		"$or": []map[string]interface{}{
			{
				"member.role": models.Developer,
			},
			{
				"member.role": models.Lead,
			},
		},
	}

	return read(rep, filter)
}

// ReadDevelopers Читает только разработчиков - member.developer
func (rep *UserDBRepository) ReadDevelopers() ([]*models.User, error) {
	filter := bson.M{
		"member.role": models.Developer,
	}

	return read(rep, filter)
}

// ReadLeadDevelopers считывает всех лид-разработчиков - member.lead
func (rep *UserDBRepository) ReadLeadDevelopers() ([]*models.User, error) {
	filter := bson.M{
		"member.role": models.Lead,
	}

	return read(rep, filter)
}

// ReadManagers считывает всех менеджеров - member.manager
func (rep *UserDBRepository) ReadManagers() ([]*models.User, error) {
	filter := bson.M{
		"$and": []map[string]interface{}{
			{
				"member.department": models.Managers,
			},
			{
				"member.role": models.Manager,
			},
		},
	}

	return read(rep, filter)
}

func read(rep *UserDBRepository, filter bson.M) ([]*models.User, error) {

	err := rep.DB.client.Ping(rep.DB.cntx, nil)

	if err != nil {
		return nil, err
	}

	collection := rep.DB.db.Collection(dbName)

	opts := options.Find()
	opts.SetLimit(100)

	cursor, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		return nil, err
	}

	var result []*models.User

	err = cursor.All(context.TODO(), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

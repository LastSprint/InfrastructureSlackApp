package repositories

import (
	"context"
	"fmt"

	models "github.com/LastSprint/InfrastructureSlackApp/models/figma"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CustomError Кастомная ошибка
type CustomError struct {
	message string
}

func (err *CustomError) Error() string {
	return err.message
}

// FigmaFilesRepository интерфейс для репозитория, который умеет работать с данными о фигме.
type FigmaFilesRepository interface {
	// ReadAllFiles считывает все файлы.
	ReadAllFiles() ([]*models.FigmaProjectFileModel, error)
	// Обновляет определенный файл.
	UpdateFile(*models.FigmaProjectFileModel) error
}

// FigmaFileDBRepository репозитория для работы с БД.
type FigmaFileDBRepository struct {
	DB *DBContext
}

const figmaDbName = "figma_files"

// ReadAllFiles считывает все файлы, которые внесены в базу данных.
func (rep *FigmaFileDBRepository) ReadAllFiles() ([]*models.FigmaProjectFileModel, error) {
	err := rep.DB.client.Ping(rep.DB.cntx, nil)

	if err != nil {
		return nil, err
	}

	collection := rep.DB.db.Collection(figmaDbName)

	opts := options.Find()
	opts.SetLimit(100)

	cursor, err := collection.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		return nil, err
	}

	var result []*models.FigmaProjectFileModel

	err = cursor.All(context.TODO(), &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateFile Обновляет определенный файл в БД.
func (rep *FigmaFileDBRepository) UpdateFile(model *models.FigmaProjectFileModel) error {
	err := rep.DB.client.Ping(rep.DB.cntx, nil)

	if err != nil {
		return err
	}

	collection := rep.DB.db.Collection(figmaDbName)

	opts := options.Find()
	opts.SetLimit(100)

	result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": model.ID}, bson.M{"$set": bson.M{"fileVersion": model.FileVersion}})

	if err != nil {
		return err
	}
	fmt.Println(result)
	if result.ModifiedCount != 1 {
		customErr := &CustomError{}
		customErr.message = "Update now work"
		return customErr
	}

	return nil
}

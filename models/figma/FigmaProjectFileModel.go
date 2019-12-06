package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FigmaProjectFileModel это модель для фигма-файла, который сязан с проектом. Эта модель находится в БД
type FigmaProjectFileModel struct {
	ID primitive.ObjectID `bson:"_id"`

	FileKey     string                 `bson:"fileKey"`
	SlackID     string                 `bson:"slackId"`
	FileVersion *FigmaFileVersionModel `bson:"fileVersion"`
}

package models

// FigmaFileVersionModel это модель версии файла Figma
type FigmaFileVersionModel struct {
	ID          string         `json:"id" bson:"fileId"`
	CreatedAt   string         `json:"created_at" bson:"createdAt"`
	Label       string         `json:"label" bson:"label"`
	Description string         `json:"description" bson:"description"`
	User        FigmaUserModel `json:"user" bson:"user"`
}

type FigmaResponseModel struct {
	Versions []FigmaFileVersionModel `json:"versions"`
}

package models

// FigmaUserModel модель пользователя в Figma
type FigmaUserModel struct {
	ID            string `json:"id" bson:"userId"`
	Handle        string `json:"handle" bson:"handle"`
	ProfileImgURL string `json:"img_url" bson:"imgUrl"`
	Email         string `json:"email" bson:"email"`
}

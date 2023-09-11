package dto

type LoginDto struct {
	Email    string `bson:"email" json:"email"  validate:"required"`
	Password string `bson:"password" json:"password" validate:"required"`
}

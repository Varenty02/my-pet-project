package config

import (
	commentmodel "MyPetProject/module/comment/model"
	likemodel "MyPetProject/module/likepost/model"
	postmodel "MyPetProject/module/posts/model"
	usermodel "MyPetProject/module/user/model"
	"gorm.io/gorm"
)

func AutoMigrateDB(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(
		&usermodel.UserCreate{},
		&commentmodel.CommentCreate{},
		&likemodel.Like{},
		&postmodel.Post{},
	)
	return db
}

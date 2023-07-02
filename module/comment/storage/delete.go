package commentstorage

import (
	"MyPetProject/commons"
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, userId, postId int) error {
	db := s.db
	if err := db.Table(commentmodel.Comment{}.TableName()).
		Where("user_id=? and post_id=?", userId, postId).
		Delete(nil).
		Error; err != nil {
		return commons.ErrDB(err)
	}
	return nil
}

package commentstorage

import (
	"MyPetProject/commons"
	commentmodel "MyPetProject/module/comment/model"
	"context"
)

func (s *sqlStore) Update(ctx context.Context, userId, postId int, data *commentmodel.CommentUpdate) (*commentmodel.CommentUpdate, error) {
	if err := s.db.Where("user_id=? and post_id=? ", userId, postId).Updates(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	return data, nil
}

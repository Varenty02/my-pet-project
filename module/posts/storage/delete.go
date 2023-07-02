package poststorage

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(postmodel.Post{}.TableName()).Where("id=?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return commons.ErrDB(err)
	}
	return nil
}

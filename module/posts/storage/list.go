package poststorage

import (
	"MyPetProject/commons"
	postmodel "MyPetProject/module/posts/model"
	"context"
	"errors"
	"strconv"
)

func (s *sqlStore) List(ctx context.Context, filter *postmodel.Filter, paging *commons.Paging, moreKey ...string) ([]postmodel.Post, error) {
	var data = []postmodel.Post{}
	db := s.db.Table(postmodel.Post{}.TableName()).Select("SUBSTRING_INDEX(content,' ',20) as content,id,name")
	if f := filter; f != nil {
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	//preload
	for i := range moreKey {
		db = db.Preload(moreKey[i])
	}
	if v := paging.Cursor; v != "" {
		if cursorId, err := strconv.Atoi(v); err != nil {
			return nil, errors.New("Cursor isn't string")
		} else {
			db = db.Where("id<?", cursorId)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
		return nil, commons.ErrDB(err)
	}
	if len(data) > 0 {
		last := data[len(data)-1]
		paging.NextCursor = strconv.Itoa(last.Id)
	}
	return data, nil
}

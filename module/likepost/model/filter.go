package likemodel

type Filter struct {
	PostId int `json:"-" form:"post_id"`
	UserId int `json:"-" form:"user_id"`
}

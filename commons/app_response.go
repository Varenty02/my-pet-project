package commons

type successRes struct {
	Data   interface{} `json:"data"`
	Filter interface{} `json:"filter,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data, filter, paging interface{}) *successRes {
	return &successRes{
		Data:   data,
		Filter: filter,
		Paging: paging,
	}
}
func SimpleResponse(data interface{}) *successRes {
	return &successRes{
		data, nil, nil,
	}
}

package model

import (
	"math"
	"reflect"
)

type Pagination struct {
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PerPage  int         `json:"per_page"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
}

func NewPagination(total int, page int, perPage int, data interface{}) Pagination {
	lastPage := math.Ceil(float64(total) / float64(perPage))
	p := Pagination{
		Total:    total,
		Page:     page,
		PerPage:  perPage,
		LastPage: int(lastPage),
		Data:     data,
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Slice && v.Len() == 0 {
		p.Data = []bool{}
	}

	return p
}

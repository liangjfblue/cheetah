package models

import "reflect"

type Page struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int         `json:"totalCount"`
	List       interface{} `json:"list"`
}

//PageUtil format list result
func PageUtil(count int, page int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}

	//防止返回null，导致接口格式不一致
	if reflect.ValueOf(list).IsNil() {
		list = []int{}
	}
	return Page{Page: page, PageSize: pageSize, TotalPage: tp, TotalCount: count, List: list}
}

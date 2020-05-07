/*
@Time : 2020/5/4 22:21
@Author : liangjiefan
*/
package models

type MenuAddRequest struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	ParentID    uint32 `json:"parentId"`
	Sequence    uint32 `json:"sequence"`
	MenuType    uint32 `json:"menuType"`
	MenuCode    string `json:"menuCode"`
	Icon        string `json:"icon"`
	OperateType string `json:"operateType"`
	IsAvailable uint32 `json:"isAvailable"`
	Remark      string `json:"remark"`
}

type MenuAddRespond struct {
	Code int32 `json:"code"`
}

type MenuDeleteRequest struct {
	Id []uint `json:"id"`
}

type MenuDeleteRespond struct {
	Code int32 `json:"code"`
}

type MenuGetRequest struct {
	Id uint `json:"id"`
}

type MenuGetRespond struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	ParentID    uint32 `json:"parentId"`
	Sequence    uint32 `json:"sequence"`
	MenuType    uint32 `json:"menuType"`
	MenuCode    string `json:"menuCode"`
	Icon        string `json:"icon"`
	OperateType string `json:"operateType"`
	IsAvailable uint32 `json:"isAvailable"`
	Remark      string `json:"remark"`
}

type MenuListRequest struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Name     string `json:"name"`
}

type Menu struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	ParentID    uint32 `json:"parentId"`
	Sequence    uint32 `json:"sequence"`
	MenuType    uint32 `json:"menuType"`
	MenuCode    string `json:"menuCode"`
	Icon        string `json:"icon"`
	OperateType string `json:"operateType"`
	IsAvailable uint32 `json:"isAvailable"`
	Remark      string `json:"remark"`
}
type MenuListRespond struct {
	Code     int32  `json:"code"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Count    int32  `json:"count"`
	Menus    []Menu `json:"menus"`
}

type MenuUpdateRequest struct {
	Id          int    `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	ParentID    uint32 `json:"parentId"`
	Sequence    uint32 `json:"sequence"`
	MenuType    uint32 `json:"menuType"`
	MenuCode    string `json:"menuCode"`
	Icon        string `json:"icon"`
	OperateType string `json:"operateType"`
	IsAvailable uint32 `json:"isAvailable"`
	Remark      string `json:"remark"`
}

type MenuUpdateRespond struct {
	Code int32 `json:"code"`
}

type MenuMenuButtonsRequest struct {
	UserId   int32  `json:"userId"`
	MenuCode string `json:"menuCode"`
}

type MenuMenuButtonsRespond struct {
	Code         int32    `json:"code"`
	Count        int32    `json:"count"`
	OperateTypes []string `json:"operateTypes"`
}

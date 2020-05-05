/*
@Time : 2020/5/4 17:07
@Author : liangjiefan
*/
package models

type RoleAddRequest struct {
	RoleName    string `json:"roleName"`
	RoleDesc    string `json:"roleDesc"`
	IsAvailable uint32 `json:"isAvailable"`
	IsAdmin     uint32 `json:"isAdmin"`
	IsBase      uint32 `json:"isBase"`
	Sequence    uint32 `json:"sequence"`
	ParentID    uint32 `json:"parentId"`
}

type RoleAddRespond struct {
	Code int32 `json:"code"`
}

type RoleDeleteRequest struct {
	Id []uint `json:"id"`
}

type RoleDeleteRespond struct {
	Code int32 `json:"code"`
}

type RoleGetRequest struct {
	Id uint `json:"id"`
}

type RoleGetRespond struct {
	Code        int32  `json:"code"`
	RoleName    string `json:"roleName"`
	RoleDesc    string `json:"roleDesc"`
	IsAvailable uint32 `json:"isAvailable"`
	IsAdmin     uint32 `json:"isAdmin"`
	IsBase      uint32 `json:"isBase"`
	Sequence    uint32 `json:"sequence"`
	ParentID    uint32 `json:"parentId"`
}

type RoleListRequest struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Name     string `json:"name"`
}

type Role struct {
	RoleName    string `json:"roleName"`
	RoleDesc    string `json:"roleDesc"`
	IsAvailable uint32 `json:"isAvailable"`
	IsAdmin     uint32 `json:"isAdmin"`
	IsBase      uint32 `json:"isBase"`
	Sequence    uint32 `json:"sequence"`
	ParentID    uint32 `json:"parentId"`
}
type RoleListRespond struct {
	Code     int32  `json:"code"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Count    int32  `json:"count"`
	Roles    []Role `json:"roles"`
}

type RoleUpdateRequest struct {
	Id          int    `json:"id"`
	RoleName    string `json:"roleName"`
	RoleDesc    string `json:"roleDesc"`
	IsAvailable uint32 `json:"isAvailable"`
	IsAdmin     uint32 `json:"isAdmin"`
	IsBase      uint32 `json:"isBase"`
	Sequence    uint32 `json:"sequence"`
	ParentID    uint32 `json:"parentId"`
}

type RoleUpdateRespond struct {
	Code int32 `json:"code"`
}

type RoleSetMenusRequest struct {
	RoleId  int32  `json:"id"`
	MenuIds []uint `json:"menuIds"`
}

type RoleSetMenusRespond struct {
	Code int32 `json:"code"`
}

type RoleAllMenusRequest struct {
	RoleId int32 `json:"id"`
}

type RoleAllMenusRespond struct {
	Code    int32  `json:"code"`
	MenuIds []uint `json:"menuIds"`
}

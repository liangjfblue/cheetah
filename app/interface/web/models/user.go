package models

type UserAddRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}

type UserAddRespond struct {
	Uid string `json:"uid"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginRespond struct {
	Token string `json:"token"`
}

type UserGetRequest struct {
	Uid string `json:"uid"`
}

type UserGetRespond struct {
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}

type UserAuthRequest struct {
	Token string `json:"token" validate:"required"`
}

type UserAuthResponse struct {
	Code int32  `json:"code"`
	UID  string `json:"uid"`
}

type UserListRequest struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Username string `json:"username"`
}

type User struct {
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}
type UserListRespond struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Count    int32  `json:"count"`
	Users    []User `json:"users"`
}

type UserUpdateRequest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}

type UserUpdateRespond struct {
	Code int32 `json:"code"`
}

type UserSetRoleRequest struct {
	UserId uint `json:"userId"`
	RoleId uint `json:"roleId"`
}

type UserSetRoleRespond struct {
	Code int32 `json:"code"`
}

type UserDeleteRequest struct {
	Id []uint `json:"id"`
}

type UserDeleteRespond struct {
	Code int32 `json:"code"`
}

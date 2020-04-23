package models

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}

type RegisterRespond struct {
	Code int32  `json:"code"`
	Uid  string `json:"uid"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespond struct {
	Code  int32  `json:"code"`
	Token string `json:"token"`
}

type GetRequest struct {
	Uid string `json:"uid"`
}

type GetRespond struct {
	Code     int32  `json:"code"`
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}

type AuthRequest struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	Code int32  `json:"code"`
	UID  string `json:"uid"`
}

type ListRequest struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Username string `json:"username"`
}

type User struct {
	Username string `json:"username"`
	Age      int32  `json:"age"`
	Addr     string `json:"addr"`
}
type ListRespond struct {
	Code     int32  `json:"code"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"pageSize"`
	Count    int32  `json:"count"`
	Users    []User `json:"users"`
}

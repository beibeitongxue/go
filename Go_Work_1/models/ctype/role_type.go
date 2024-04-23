package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin   Role = 1 //管理员
	PermissionUser    Role = 2 //普通用户
	PermissionVip     Role = 3 //Vip用户
	PermissionVisitor Role = 4 //游客
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s Role) String() string {
	var str string
	switch s {
	case PermissionAdmin:
		str = "管理员"
	case PermissionUser:
		str = "用户"
	case PermissionVip:
		str = "vip"
	case PermissionVisitor:
		str = "游客"
	default:
		str = "其他"

	}
	return str
}

package res

type ErrCode int

func (e ErrCode) Error() string {
	//TODO implement me
	panic("implement me")
}

const (
	SettingsError   ErrCode = 1001
	ArgumentError   ErrCode = 1002
	UpdateFailed    ErrCode = 1003
	NodeNotFound    ErrCode = 1004
	CreatedError    ErrCode = 1005
	UserNameError   ErrCode = 1006
	TokenslideError ErrCode = 1007
	UserNotExit     ErrCode = 1008
	EmailBlindError ErrCode = 1009
	EmailExit       ErrCode = 1010
	PasswordError   ErrCode = 1011
	NodeExit        ErrCode = 1012
)

var (
	ErrMap = map[ErrCode]string{
		SettingsError:   "系统错误",
		ArgumentError:   "参数错误",
		UpdateFailed:    "修改失败",
		NodeNotFound:    "节点不存在",
		CreatedError:    "创建失败",
		UserNameError:   "用户名或密码错误",
		UserNotExit:     "用户不存在",
		EmailBlindError: "邮箱绑定失败",
		EmailExit:       "邮箱重复",
		PasswordError:   "密码不一致",
	}
)

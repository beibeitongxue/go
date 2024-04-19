package res

type ErrCode int

const (
	SettingsError ErrCode = 1001
	ArgumentError ErrCode = 1002
)

var (
	ErrMap = map[ErrCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)

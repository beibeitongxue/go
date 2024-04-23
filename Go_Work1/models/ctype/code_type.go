package ctype

import "encoding/json"

type Node int

const (
	TrailNodes   Node = 1 //试用节点
	VipNodes     Node = 2 //付费节点
	PrivateNodes Node = 3 //私有节点

)

func (s Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
func (s Node) String() string {
	var str string
	switch s {
	case TrailNodes:
		str = "试用节点"
	case VipNodes:
		str = "付费节点"
	case PrivateNodes:
		str = "私有节点"
	default:
		str = "其他"

	}
	return str
}

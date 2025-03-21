package appTypes

import "encoding/json"

// Storage 图片存储类型
type Storage int

const (
	Local Storage = iota //本地存储
	Qiniu                //七牛云存储
)

// MarshalJSON 实现了 json.Marshaler 接口
func (s Storage) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

// String 方法返回 Storage 的字符串表示
func (s Storage) String() string {
	var str string
	switch s {
	case Local:
		str = "本地"
	case Qiniu:
		str = "七牛云"
	default:
		str = "未知存储"
	}
	return str
}

// ToStorage函数将字符串转换为Storage
func ToStorage(str string) Storage {
	switch str {
	case "本地":
		return Local
	case "七牛云":
		return Qiniu
	default:
		return -1
	}
}

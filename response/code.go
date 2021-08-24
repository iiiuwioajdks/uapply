package response

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
}

func (c ResCode) Msg() string {
	return CodeMsgMap[c]
}

package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeServerBusy
)

var codeMsgMaP = map[ResCode]string{
	CodeSuccess:    "success",
	CodeServerBusy: "server busy",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMaP[c]
	if !ok {
		msg = codeMsgMaP[CodeServerBusy]
	}
	return msg
}

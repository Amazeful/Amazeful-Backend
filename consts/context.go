package consts

type ContextKey string

const (
	CtxSession ContextKey = "session"
	CtxChannel ContextKey = "channel"
	CtxUser    ContextKey = "user"
	CtxCommand ContextKey = "command"
)

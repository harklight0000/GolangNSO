package constants

//go:generate go run ../generate/gen_string_constants.go -pkg=constants -type=HookType
type HookType int

const (
	HOOK_NET_WORK_START HookType = iota
	HOOK_NEW_CONNECTION_ENTER
	HOOK_NET_WORK_STOP
	HOOK_APP_START
	HOOK_APP_TERMINATE
	HOOK_APP_END
)

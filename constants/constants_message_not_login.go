package constants

//go:generate go run ../generate/gen_string_constants.go -pkg=constants -type=MessageNotLogin
type MessageNotLogin int

const (
	LOGIN0           MessageNotLogin = 0
	REGISTER0        MessageNotLogin = 1
	SEND_SMS         MessageNotLogin = -124
	CLEAR_RMS        MessageNotLogin = 2
	CLIENT_INFO      MessageNotLogin = -125
	LOGIN            MessageNotLogin = -127
	FORGET_PASS      MessageNotLogin = -122
	FORGET_PASS_IMEI MessageNotLogin = -121
	REGISTER         MessageNotLogin = -126
	REGISTER_IMEI    MessageNotLogin = -123
	DOI_MAT_KHAU     MessageNotLogin = 120
)

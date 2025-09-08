package constants

//go:generate go run ../generate/gen_string_constants.go -pkg=constants -type=MessageNotMap
//-28
type MessageNotMap int

const (
	GPS                                  MessageNotMap = -59
	CLEAR_TASK                                         = -98
	INPUT_CARD                                         = -99
	CHANGE_NAME                                        = -97
	REQUEST_ICON                                       = -115
	UPDATE_PK                                          = -117
	CREATE_CLAN                                        = -96
	CLAN_CHANGE_TYPE                                   = -94
	REQUEST_CLAN_INFO                                  = -113
	CLAN_MOVEOUT_MEM                                   = -93
	CLAN_OUT                                           = -92
	CLAN_UP_LEVEL                                      = -91
	INVITE_CLANDUN                                     = -87
	REQUEST_CLAN_LOG                                   = -114
	OPEN_CLAN_ITEM                                     = -62
	ITEM_SPLIT                                         = -85
	CHIENTRUONG_INFO                                   = -81
	CHANGE_BG_ID                                       = -77
	MOI_GTC                                            = -70
	LAT_HINH                                           = -72
	CONVERT_UPGRADE                                    = -88
	REQUEST_CLAN_MEMBER                                = -112
	REQUEST_CLAN_ITEM                                  = -111
	UPDATE_OUT_CLAN                                    = -116
	INPUT_COIN_CLAN                                    = -90
	OUTPUT_COIN_CLAN                                   = -89
	NOT_USEACC                                         = -86
	ME_LOAD_ACTIVE                                     = -106
	POINT_PB                                           = -84
	REVIEW_CT                                          = -80
	REVIEW_PB                                          = -83
	CLAN_CHANGE_ALERT                                  = -95
	SELECT_PLAYER                                      = -126
	UPDATE_VERSION                                     = -123
	UPDATE_DATA                                        = -122
	UPDATE_MAP                                         = -121
	UPDATE_SKILL                                       = -120
	UPDATE_ITEM                                        = -119
	REQUEST_MOB_TEMPLATE                               = -108
	REQUEST_MAP_TEMPLATE                               = -109
	OAN_HON                                            = -67
	OAN_HON1                                           = -66
	CLIENT_OK                                          = -101
	CLEAR_ACC_PROTECT                                  = -102
	OPEN_LOCK_ACC_PROTECT                              = -103
	UPDATE_LOCK_ACC_PROTECT                            = -104
	ACTIVE_ACC_PROTECT                                 = -105
	REQUEST_NPC_TEMPLATE                               = -107
	REQUEST_SKILL                                      = -110
	ACCEPT_GT_CHIEN                                    = -68
	MOI_TATCA_GTC                                      = -69
	REWARD_PB                                          = -82
	CHAT_ADMIN                                         = -78
	REWARD_CT                                          = -79
	CLAN_SEND_ITEM                                     = -61
	CLAN_USE_ITEM                                      = -60
	REQUEST_IMG_EFF_AUTO_SEND_OR_CAPTCHA               = 122
)

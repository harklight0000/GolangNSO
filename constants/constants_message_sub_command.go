package constants

//go:generate go run ../generate/gen_string_constants.go -pkg=constants -type=MessageSubCommand
type MessageSubCommand int

// -30
const (
	LOAD_THU_CUOI                         MessageSubCommand = -54
	PLAYER_LOAD_GLOVE                                       = -55
	PLAYER_LOAD_AO_CHOANG                                   = -56
	CALL_EFFECT_BALL_1                                      = -57
	CALL_EFFECT_BALL                                        = -58
	REFRESH_HP                                              = -59
	CLAN_ACCEPT_PLEASE                                      = -60
	CLAN_PLEASE                                             = -61
	CLAN_ACCEPT_INVITE                                      = -62
	CLAN_INVITE                                             = -63
	PLAYER_LOAD_MAT_NA                                      = -64
	SEND_SKILL                                              = -65
	SAVE_RMS                                                = -67
	PLAYER_LOAD_THU_NUOI                                    = -68
	ME_LOAD_THU_NUOI                                        = -69
	ADMIN_MOVE                                              = -70
	ME_UP_GOLD                                              = -71
	ME_LOAD_GOLD                                            = -72
	CALL_EFFECT_MOB                                         = -73
	SHOW_WAIT                                               = -74
	ITEM_BOX_CLEAR                                          = -75
	LOCK_PARTY                                              = -76
	FIND_PARTY                                              = -77
	CALL_EFFECT_ME                                          = -78
	BUFF_LIVE                                               = -79
	ITEM_BODY_CLEAR                                         = -80
	ME_UPDATE_PK                                            = -81
	ENEMIES_REMOVE                                          = -82
	FRIEND_REMOVE                                           = -83
	REQUEST_ENEMIES                                         = -84
	REQUEST_FRIEND                                          = -85
	MOVE_MEMBER                                             = -86
	CHANGE_TEAMLEADER                                       = -87
	CREATE_PARTY                                            = -88
	END_WAIT                                                = -89
	TASK_FOLLOW_FAIL                                        = -90
	UPDATE_BAG_EXPAND                                       = -91
	UPDATE_TYPE_PK                                          = -92
	CHANGE_TYPE_PK                                          = -93
	NPC_PLAYER_UPDATE                                       = -94
	MAP_TIME                                                = -95
	PLAYER_REMOVE_EFFECT                                    = -96
	PLAYER_EDIT_EFFECT                                      = -97
	PLAYER_ADD_EFFECT                                       = -98
	ME_REMOVE_EFFECT                                        = -99
	ME_EDIT_EFFECT                                          = -100
	ME_ADD_EFFECT                                           = -101
	USE_BOOK_SKILL                                          = -102
	REQUEST_ITEM                                            = -103
	BOX_COIN_OUT                                            = -104
	BOX_COIN_IN                                             = -105
	BOX_SORT                                                = -106
	BAG_SORT                                                = -107
	SKILL_UP                                                = -108
	POTENTIAL_UP                                            = -109
	PLAYER_LOAD_LIVE                                        = -110
	PLAYER_LOAD_HP                                          = -111
	PLAYER_LOAD_REFRESH_HP_EFF5_BUFF_FULL                   = -112
	PLAYER_LOAD_QUAN                                        = -113
	PLAYER_LOAD_AO                                          = -116
	PLAYER_LOAD_VUKHI                                       = -117
	PLAYER_LOAD_INFO                                        = -119
	PLAYER_LOAD_ALL                                         = -120
	ME_LOAD_MP                                              = -121
	ME_LOAD_HP                                              = -122
	ME_LOAD_INFO                                            = -123
	ME_LOAD_LEVEL                                           = -124
	UPDATE_INFO_ME                                          = 115
	ME_LOAD_SKILL                                           = -125
	ME_LOAD_CLASS                                           = -126
	ME_LOAD_ALL                                             = -127
	PLAYER_LOAD_LEVEL                                       = -128
)

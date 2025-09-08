package controllers

import (
	"fmt"
	. "nso/constants"
	"nso/core"
	"nso/logging"
	"nso/networking"
	"nso/objects"
	"nso/utils"
)

func (this *Controller) MessageNotLogin(session *networking.Session, message *core.Message, user *objects.User) error {
	cmd := MessageNotLogin(message.ReadByte())
	logging.Logger.Info(fmt.Sprintf("Enter Message not login command id = %d, name = %s", cmd, cmd.String()))
	switch cmd {
	case LOGIN:
		utils.Bench(func() {
			this.Login(session, message)
		})

	case CLIENT_INFO:
		return this.SetClientInfo(session, message)
	default:
		logging.Logger.Info(fmt.Sprintf("Unknown command id = %d, name = %s", cmd, cmd.String()))
	}
	return nil
}

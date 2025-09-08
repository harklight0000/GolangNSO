package clandun

import (
	. "nso/ainterfaces"
	"nso/utils"
)

type ClanDunManager struct {
	baseClanID utils.AtomicInteger
	clanDun    map[int]*ClanDun
	appCtx     IAppContext
}

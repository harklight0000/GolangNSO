package objects

import (
	. "nso/ainterfaces"
	"sync"
)

type ClansManager struct {
	lock  sync.RWMutex
	Clans map[string]*ClanManager
}

func NewClansManager() *ClansManager {
	return &ClansManager{}
}

func (this *ClansManager) GetClanByName(clanName string) IClanManager {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.Clans[clanName]
}

func (this *ClansManager) CreateClan(user IUser, name string) (IClanManager, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	// TODO
	return nil, nil
}

func (this *ClansManager) AddClan(clanManager IClanManager) {
	//TODO implement me
	panic("implement me")
}

func (this *ClansManager) Dissolution(clanName string) error {
	//TODO implement me
	panic("implement me")
}

func (this *ClansManager) Close() error {
	//TODO implement me
	panic("implement me")
}

package ainterfaces

type IClansManager interface {
	GetClanByName(clanName string) IClanManager
	CreateClan(user IUser, name string) (IClanManager, error)
	AddClan(clanManager IClanManager)
	Dissolution(clanName string) error
	Close() error
}

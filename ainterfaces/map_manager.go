package ainterfaces

type IMapManager interface {
	GetMapByID(id int) IMap
	Enter(user IUser, dst IMap) error
}

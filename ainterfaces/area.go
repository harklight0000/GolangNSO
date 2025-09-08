package ainterfaces

type IArea interface {
	IGameObject
	ISender
	RefreshBoss()
	Enter(user IUser)
	Enter0(user IUser)
	IsFree() bool
}

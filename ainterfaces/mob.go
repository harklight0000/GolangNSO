package ainterfaces

type IMob interface {
	IGameObject
	IncHp(num int)
	ClearFight()
	Fight(ninjaId int, dame int)
	CheckFight(ninjaId int) bool
}

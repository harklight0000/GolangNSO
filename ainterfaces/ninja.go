package ainterfaces

type INinja interface {
	ISender
	UpdateExpNormal(exp int64)
	UpdateExpUseMulti(exp int64)
	UpdatePBPoint(point int16)
	GetName() string
	GetID() int
	NClass() byte
	GetLevel() int
	GetParty() IParty
	SendYellowMessage(s string)
	GetX() int16
	GetY() int16
}

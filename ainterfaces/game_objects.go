package ainterfaces

type IParseData interface {
	ParseData() error
}

type IUpdate interface {
	UpdateAsync() error
	Update() error
}

type IGameObject interface {
	Awake() error
	Start() error
	IUpdate
	Close() error
}

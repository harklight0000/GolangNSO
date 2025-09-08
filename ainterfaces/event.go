package ainterfaces

type Event struct {
	Func  func(area IArea) func() error
	Func2 func(area IArea) func()
}

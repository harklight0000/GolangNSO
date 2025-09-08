package ainterfaces

type IUserManager interface {
	GetUser(id int) (user IUser, err error)
	GetUsers() (users []IUser)
	AddUser(user IUser) (err error)
	RemoveUser(id int) (err error)
	AddSession(session ISession) (err error)
	RemoveSession(id int) (err error)
	Size() int
}

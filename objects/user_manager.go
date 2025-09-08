package objects

import (
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	. "nso/ainterfaces"
	"nso/logging"
	"nso/utils"
	"sync"
)

var (
	ErrUserAlreadyExists    = eris.New("User already exists")
	ErrSessionAlreadyExists = eris.New("Session already exists")
)

type UserManager struct {
	lUsers sync.RWMutex
	users  map[int]IUser

	lSessions sync.RWMutex
	lNinjas   sync.RWMutex
	sessions  map[int]ISession

	ninjas map[int]INinja
	ctx    IAppContext
}

func (this *UserManager) Size() int {
	this.lUsers.RLock()
	defer this.lUsers.RUnlock()
	return len(this.users)
}

func NewUserManager(ctx IAppContext) *UserManager {
	this := &UserManager{}
	this.users = make(map[int]IUser)
	this.sessions = make(map[int]ISession)
	this.ninjas = make(map[int]INinja)
	this.ctx = ctx
	return this
}

func (this *UserManager) GetUser(id int) (user IUser, err error) {
	this.lUsers.RLock()
	defer this.lUsers.RUnlock()
	user, ok := this.users[id]
	if !ok {
		return nil, eris.New("User not found")
	}
	return user, nil
}

func (this *UserManager) GetUsers() (users []IUser) {
	this.lUsers.RLock()
	defer this.lUsers.RUnlock()
	for _, user := range this.users {
		users = append(users, user)
	}
	return users
}

func (this *UserManager) AddUser(user IUser) (err error) {
	this.lUsers.Lock()
	defer this.lUsers.Unlock()
	this.users[user.GetID()] = user
	nj := this.users[user.GetID()].GetNinja()
	this.lNinjas.Lock()
	defer this.lNinjas.Unlock()
	if !utils.IsNil(nj) {
		this.ninjas[nj.GetID()] = nj
	}
	clone := this.users[user.GetID()].GetClone()
	if !utils.IsNil(clone) {
		this.ninjas[clone.GetID()] = clone
	}
	return
}

func (this *UserManager) RemoveUser(id int) (err error) {
	this.lUsers.Lock()
	defer this.lUsers.Unlock()
	user, ok := this.users[id]
	delete(this.users, id)
	if ok && user != nil {
		this.lNinjas.Lock()
		defer this.lNinjas.Unlock()
		nj := user.GetNinja()
		if !utils.IsNil(nj) {
			delete(this.ninjas, nj.GetID())
		}
		clone := user.GetClone()
		if !utils.IsNil(clone) {
			delete(this.ninjas, clone.GetID())
		}
	}
	return
}

func (this *UserManager) AddSession(session ISession) (err error) {
	this.lSessions.Lock()
	defer this.lSessions.Unlock()
	this.sessions[session.GetID()] = session
	return
}

func (this *UserManager) RemoveSession(id int) (err error) {
	this.lSessions.Lock()
	ss := this.sessions[id]
	delete(this.sessions, id)
	this.lSessions.Unlock()
	if ss != nil {
		user, ok := ss.GetUser().(*User)
		if ok && user != nil {
			err = this.RemoveUser(user.ID)
		}
	}
	logging.Logger.Info("Remove session", zap.Int("id", id))
	return
}

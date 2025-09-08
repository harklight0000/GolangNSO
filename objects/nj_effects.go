package objects

import "sync"

type NjEffects struct {
	effects map[int]*Effect
	effLock sync.RWMutex
}

func (this *NjEffects) HasEffect(id int) bool {
	this.effLock.RLock()
	defer this.effLock.RUnlock()
	return this.effects[id] != nil
}

func (this *NjEffects) AddEffect(effect *Effect) {
	this.effLock.Lock()
	defer this.effLock.Unlock()
	this.effects[effect.ID] = effect
	// TODO
}

func (this *NjEffects) RemoveEffect(id int) {
	this.effLock.Lock()
	defer this.effLock.Unlock()
	delete(this.effects, id)
}

func (this *NjEffects) GetEffectByType(t int) *Effect {
	this.effLock.RLock()
	defer this.effLock.RUnlock()
	for _, v := range this.effects {
		if v.Type == t {
			return v
		}
	}
	return nil
}

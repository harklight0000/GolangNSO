package objects

import (
	. "nso/ainterfaces"
	"nso/core"
	"nso/utils"
)

var baseId = utils.NewAtomicInteger(0)

type Party struct {
	isLock     bool
	id         int
	cave       *Cave
	master     *Ninja
	ninjas     []*Ninja
	invitation map[int]bool
	battle     IBattle
}

func (this *Party) AcceptParty(ninja INinja) {
	if !this.invitation[ninja.GetID()] {
		ninja.SendYellowMessage("Bạn không có lời mời")
		return
	}
	this.invitation[ninja.GetID()] = false
	this.ninjas = append(this.ninjas, ninja.(*Ninja))
	this.SendText(ninja.GetName() + " đã tham gia nhóm")
	delete(this.invitation, ninja.GetID())
	this.RefreshTeam()
}

func (this *Party) RefreshTeam() {
	RefreshTeam(this)
}

func (this *Party) SendText(text string) {
	for _, v := range this.ninjas {
		v.SendYellowMessage(text)
	}
}

func (this *Party) ChatMessage(ninja INinja, text string) {
	ChatParty(this, ninja.GetName(), text)
}

func NewParty(master *Ninja) *Party {
	this := &Party{id: baseId.IncAndGet(), master: master}
	this.ninjas = append(this.ninjas, master)
	this.invitation = make(map[int]bool)
	return this
}

func (this *Party) GetMaster() INinja {
	return this.master
}

func (this *Party) IsLocked() bool {
	return this.isLock
}

func (this *Party) GetCave() ICave {
	return this.cave
}

func (this *Party) GetID() int {
	return this.id
}

func (this *Party) GetMapID() int {
	return int(this.master.MapID)
}

func (this *Party) GetNinjaByID(id int) INinja {
	for _, v := range this.ninjas {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func (this *Party) OpenCave(cave ICave) {
	this.cave = cave.(*Cave)
	this.SendText(this.master.Name + " đã mở Cửa hang động")
}

func (this *Party) GetNinjas() []INinja {
	var ninjas []INinja
	for _, v := range this.ninjas {
		ninjas = append(ninjas, v)
	}
	return ninjas
}

func (this *Party) GetBattle() IBattle {
	return this.battle
}

func (this *Party) SetBattle(battle IBattle) {
	this.battle = battle
}

func (this *Party) UpdateEffect(effect IEffect) {
	for _, v := range this.ninjas {
		ef := effect.(*Effect)
		v.AddEffect(ef.ID, ef.TimeStart, ef.TimeLength, ef.Param)
	}
}

func (this *Party) GetName() string {
	if this.master == nil {
		return ""
	}
	return this.master.Name
}

func (this *Party) ChangeTeamLeader(index int) {
	if index < 0 || index >= len(this.ninjas) {
		return
	}
	this.master = this.ninjas[index]
	this.SendText(this.GetName() + " đã được chọn làm nhóm trưởng")
	this.RefreshTeam()
}

func (this *Party) AddToParty(ninja INinja) {
	if len(this.ninjas) >= 6 {
		this.master.SendYellowMessage("Số lượng thành viên đã tối đa")
		return
	}
	party, ok := ninja.GetParty().(*Party)
	if ok && party != nil {
		this.master.SendYellowMessage("Đối phương đã có nhóm")
		return
	}
	this.invitation[ninja.GetID()] = true
	SendInvitationParty(this, ninja)
}
func (this *Party) RemoveMember(ninja INinja) {
	this.removeMember(ninja, "bị đuổi khỏi nhóm")
}
func (this *Party) removeMember(ninja INinja, reason string) {
	var removedNinja INinja
	for i, v := range this.ninjas {
		if v == ninja {
			removedNinja = v
			this.ninjas = append(this.ninjas[:i], this.ninjas[i+1:]...)
			break
		}
	}
	if !utils.IsNil(removedNinja) {
		removedNinja.SendYellowMessage("Bạn đã " + reason)
		this.SendText(removedNinja.GetName() + " đã " + reason)
		this.RefreshTeam()
	}
}

func (this *Party) ExitParty(ninja INinja) {
	removed := this.GetNinjaByID(ninja.GetID())
	if utils.IsNil(removed) {
		ninja.SendYellowMessage("Bạn không ở trong nhóm")
		return
	}
	this.removeMember(removed, "rời khởi nhóm")
}

func (this *Party) EnterArea(area IArea) {
	_area := area.(*Area)
	for _, v := range this.ninjas {
		_area.Enter0(v.User)
	}
}

func (this *Party) ClearBattle() {

}

func (this *Party) SendToAll(m *core.Message) {
	for _, v := range this.ninjas {
		v.SendMessage(m)
	}
}

func (this *Party) Size() byte {
	return byte(len(this.ninjas))
}

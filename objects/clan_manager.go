package objects

import (
	"encoding/json"
	"fmt"
	"github.com/rotisserie/eris"
	. "nso/ainterfaces"
	"nso/core"
	"nso/entity"
)

type ClanManager struct {
	*entity.ClanEntity
	Members []*ClanMember
	Items   []*Item
	AppCtx  IAppContext
}

func NewClanManager(clanEntity *entity.ClanEntity, appCtx IAppContext) *ClanManager {
	this := &ClanManager{}
	this.ClanEntity = clanEntity
	this.AppCtx = appCtx
	return this
}

func (this *ClanManager) GetName() string {
	return this.Name
}

func (this *ClanManager) UpdateCoin(coin int) {
	if coin > 0 {
		this.Coin += coin
	}
}

func (this *ClanManager) GetTocPho() ClanMember {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetTocTruong() ClanMember {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetMembers() []ClanMember {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetElders() []ClanMember {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetMemberByID(cID int) ClanMember {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetMemberByName(name string) ClanMember {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetMemberMax() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetExpNext() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetFreeCoin() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetCoinOpen() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetCoinUp() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SendMessage(message *core.Message) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) PayFeeClan() {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) UpExp(exp int) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) AddItem(item IItem) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) RemoveItem(item IItem, quantity int) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) Chat(user IUser, message string) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) ChangeTypeClan(user IUser, m core.Message) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) OpenItemLevel(user IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SendClanItem(user IUser, m core.Message) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SetAlert(user IUser, m core.Message) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) MoveOutClan(user IUser, m core.Message) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) CreateMessage(m string) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) OutClan(user IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) ClanUpLevel(user IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) InputCoinClan(user IUser, m core.Message) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) InformAll(message string) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) WriteLog(name string, num int, number int, date string) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) LogClan(p IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) RequestClanInfo(user IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) RequestClanMember(user IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) RequestClanItem(user IUser) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetLevel() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SetLevel(level int) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) InviteToDun(name string) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SetClanTerritory(clanTerritory interface{}) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SetClanBattle(clanBattle interface{}) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetClanTerritory() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetClanBattle() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SetClanBattleData(clanBattleData interface{}) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetClanBattleData() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) RestoreClanBattle() bool {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetCurrentThanThu() interface{} {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) GetThanThuIndex() int {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) SetThanThuIndex(index int) {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) Close() error {
	//TODO implement me
	panic("implement me")
}

func (this *ClanManager) ParseData() (error1 error) {
	defer func() {
		if err := recover(); err != nil {
			error1 = eris.New(fmt.Sprintf("Error occur when parse data: %s", err))
		}
	}()
	var members [][]interface{}
	err := json.Unmarshal([]byte(this.ClanEntity.Members), &members)
	if err != nil {
		return eris.Wrap(err, "Error when parse clan members")
	}
	for _, member := range members {
		this.Members = append(this.Members, &ClanMember{
			CharID:        cast[int](member[0]),
			CName:         member[1].(string),
			ClanName:      member[2].(string),
			TypeClan:      cast[int8](member[3]),
			CLevel:        cast[int](member[4]),
			NClass:        cast[byte](member[5]),
			PointClan:     cast[int](member[6]),
			PointClanWeek: cast[int](member[7]),
		})
	}
	return nil
}

func (this *ClanManager) GetItems() []IItem {
	items := make([]IItem, len(this.Items))
	for i, item := range this.Items {
		items[i] = item
	}
	return items
}

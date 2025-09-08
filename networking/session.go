package networking

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"io"
	"net"
	. "nso/ainterfaces"
	. "nso/constants"
	"nso/core"
	"nso/errs"
	"nso/logging"
	. "nso/utils"
	"strings"
)

var baseSessionID AtomicInteger

type Sbyte = int8

type ClientInfo struct {
	Type       byte
	ZoomLevel  byte
	IsGps      bool
	Width      int
	Height     int
	IsQwerty   bool
	IsTouch    bool
	Platform   string
	LanguageID byte
	Provider   int
	Agent      string
}

func NewSession(conn net.Conn, appCtx IAppContext) *Session {
	this := &Session{Conn: conn}
	this.ID = baseSessionID.IncAndGet()
	this.Ctx, this.cancel = context.WithCancel(appCtx.GetContext())
	this.dis = core.NewDataInputStream(conn)
	this.dos = core.NewDataOutputStream(conn)
	this.SendData = make(chan *core.Message)
	this.Connected = true
	this.AppCtx = appCtx
	this.controller = appCtx.GetController()
	this.ClientIP = conn.RemoteAddr().String()
	keys := "Test"
	this.keys = make([]Sbyte, len(keys))
	for i := 0; i < len(keys); i++ {
		this.keys[i] = Sbyte(keys[i])
	}
	return this
}

type Session struct {
	ID              int
	Connected       bool
	GetKeyComplete  bool
	SendData        chan *core.Message
	CurR            int8
	CurW            int8
	Conn            net.Conn
	dis             *core.DataInputStream
	dos             *core.DataOutputStream
	ZoomLevel       byte
	Ctx             context.Context
	cancel          context.CancelFunc
	LastTimeReceive int64
	Version         string
	ClientIP        string
	keys            []Sbyte
	User            IUser
	AppCtx          IAppContext
	controller      IController
	ClientInfo
	Login bool
}

func (this *Session) GetID() int {
	return this.ID
}
func (s Session) IsNew() bool {
	if strings.HasPrefix(s.Version, "2") {
		return true
	}
	return false
}

func (this *Session) GetUser() IUser {
	return this.User
}

func (this *Session) doSendMessage(m *core.Message) {
	data := m.GetData()
	if data != nil {
		b := m.Command
		this.logSend(m)
		size := len(data)
		if size > 65535 {
			b = -32
		}
		if this.GetKeyComplete {
			this.dos.WriteSByte(this.writeKey(int8(b)))
		} else {
			this.dos.WriteIByte(int(b))
		}
		if b == -32 {
			b = m.Command
			if this.GetKeyComplete {
				this.dos.WriteSByte(this.writeKey(int8(b)))
			} else {
				this.dos.WriteIByte(int(b))
			}
			b1 := this.writeKey(int8(size >> 24))
			b2 := this.writeKey(int8(size >> 16))
			b3 := this.writeKey(int8(size >> 8))
			b4 := this.writeKey(int8(size & 0xFF))
			this.dos.WriteSByte(b1)
			this.dos.WriteSByte(b2)
			this.dos.WriteSByte(b3)
			this.dos.WriteSByte(b4)
		} else if this.GetKeyComplete {
			this.dos.WriteSByte(this.writeKey(int8(size >> 8)))
			this.dos.WriteSByte(this.writeKey(int8(size & 0xFF)))
		} else {
			this.dos.WriteIByte(size & 0xFF00)
			this.dos.WriteIByte(size & 0xFF)
		}
		if this.GetKeyComplete {
			for i := 0; i < size; i++ {
				this.dos.WriteSByte(this.writeKey(int8(data[i])))
			}
		} else {
			for i := 0; i < size; i++ {
				this.dos.WriteIByte(int(data[i]))
			}
		}
	}
}

func (this *Session) SendMessage(message *core.Message) {
	this.SendData <- message
}

func (this *Session) ReadMessage() *core.Message {
	cmd := this.dis.ReadByte()
	if cmd != -27 {
		cmd = this.readKey(cmd)
	}
	var size int
	if cmd != -27 {
		b1 := this.dis.ReadByte()
		b2 := this.dis.ReadByte()
		size = int(this.readKey(b1))<<8 | (int(this.readKey(b2)) & 0xFF)
	} else {
		size = int(this.dis.ReadShort())
	}
	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = Byte(this.dis.ReadByte())
	}
	if cmd != -27 {
		for i := 0; i < size; i++ {
			data[i] = Byte(this.readKey(int8(data[i])))
		}
	}
	return core.NewReaderMessage(Command(cmd), bytes.NewBuffer(data))
}

func (this *Session) Update() {
	go this.send()
	go this.receive()
}

func (this *Session) Disconnect() {
	defer func() {
		if this.User != nil {
			switch this.User.(type) {
			case ISync:
				err := this.User.(ISync).Sync()
				if err != nil {
					logging.Logger.Error("Error when sync ninja", zap.Error(err))
				}
			}
			this.User.Leave()
			err := this.AppCtx.GetUserManager().RemoveUser(this.User.GetID())
			if err != nil {
				logging.Logger.Error("Error when remove user", zap.Error(err))
			}
		}
	}()
	logging.Logger.Info("Disconnecting session", zap.Int("session", this.ID))
	this.Connected = false
	this.cancel()
}

func (this *Session) send() {
	defer func() {
		if err := recover(); err != nil {
			logging.Logger.Error("Error in send: " + fmt.Sprint(err))
		}
	}()
	for this.Connected {
		select {
		case <-this.Ctx.Done():
			logging.Logger.Info("Session send done", zap.Int("session", this.ID))
			return
		case m := <-this.SendData:
			this.doSendMessage(m)
		}
	}
	logging.Logger.Info("Session " + this.ClientIP + " stop sending")
}

func (this *Session) receive() {
	defer func() {
		if err := recover(); err != nil {
			err, ok := err.(error)
			if ok {
				if errors.Is(err, io.EOF) {
					logging.Logger.Info("EOF in receive")
					this.Disconnect()
				}
			} else {
				logging.Logger.Error("Session " + this.ClientIP + " receive error: " + fmt.Sprint(err))
			}
		}
	}()
	for this.Connected {
		select {
		case <-this.Ctx.Done():
			logging.Logger.Info("Close read goroutine")
			return
		default:
			message := this.ReadMessage()
			err := this.controller.OnMessage(this, message)
			if err != nil {
				if e, ok := err.(*errs.ErrNextMap); ok {
					this.SendServerDialog(e.Error())
				} else {
					logging.Logger.Error(eris.ToString(err, true))
				}
			}
		}
	}
}

func (this *Session) writeKey(b int8) Sbyte {
	c := this.CurW
	this.CurW++
	r := (Byte(this.keys[c]) & 0xFF) ^ (Byte(b) & 0xFF)
	if int(this.CurW) >= len(this.keys) {
		this.CurW %= int8(len(this.keys))
	}
	return Sbyte(r)
}

func (this *Session) readKey(b Sbyte) Sbyte {
	r := this.CurR
	this.CurR++
	i := (Byte(this.keys[r]) & 0xFF) ^ (Byte(b) & 0xFF)
	if int(this.CurR) >= len(this.keys) {
		this.CurR %= int8(len(this.keys))
	}
	return Sbyte(i)
}

func (this *Session) SendKey() {
	m := core.NewMessage(-27)
	m.WriteIByte(len(this.keys))
	m.WriteSByte(this.keys[0])
	for i := 1; i < len(this.keys); i++ {
		m.WriteSByte(this.keys[i] ^ this.keys[i-1])
	}
	this.doSendMessage(m)
	this.GetKeyComplete = true
}

func (this *Session) UpdateVersion() {
	m := core.NewMessage(NOT_MAP)
	m.WriteSByte(UPDATE_VERSION)
	m.WriteByte(VsData)
	m.WriteByte(VsMap)
	m.WriteByte(VsSkill)
	m.WriteByte(VsItem)
	m.WriteByte(0)
	m.WriteByte(0)
	m.WriteByte(0)
	this.SendMessage(m)
}

func (this *Session) SendServerDialog(text string) {
	m := core.NewMessage(SERVER_DIALOG)
	m.WriteUTF(text)
	this.SendMessage(m)
}

func (this *Session) logSend(m *core.Message) {
	b := m.Command
	data := m.GetData()
	logging.Logger.Info(fmt.Sprintf("Sending command = %d, name=%s", b, b.String()))
	if b == NOT_MAP {
		cmd := MessageNotMap(int8(data[0]))
		logging.Logger.Info(fmt.Sprintf("\t with message not map= %d, name=%s size %d", cmd, cmd.String(), len(data)))

	} else if b == NOT_LOGIN {
		cmd := MessageNotLogin(int8(data[0]))
		logging.Logger.Info(fmt.Sprintf("\t with message not login= %d, name=%s", cmd, cmd.String()))
	} else if b == MESSAGE_SUB_COMMAND {
		cmd := MessageSubCommand(int8(data[0]))
		logging.Logger.Info(fmt.Sprintf("\t with message sub command= %d, name=%s", cmd, cmd.String()))
	}
}

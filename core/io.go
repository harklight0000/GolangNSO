package core

import (
	"bytes"
	"io"
	"nso/constants"
)

type Message struct {
	Command constants.Command
	InputStream
	OutputStream
	buffer bytes.Buffer
}

func (m *Message) GetData() []byte {
	data := m.buffer.Bytes()
	return data
}

func NewReaderMessage(command constants.Command, reader io.Reader) *Message {
	var msg = &Message{Command: command}
	msg.InputStream = NewDataInputStream(reader)
	return msg
}

// NewMessage is for writer
func NewMessage(command constants.Command) *Message {
	var msg = &Message{Command: command, buffer: bytes.Buffer{}}
	msg.OutputStream = NewDataOutputStream(&msg.buffer)
	return msg
}

// MessageNotMap command -28
func MessageNotMap(command constants.MessageNotMap) *Message {
	var msg = NewMessage(constants.NOT_MAP)
	msg.WriteIByte(int(command))
	return msg
}

// MessageSubCommand command -30
func MessageSubCommand(command constants.MessageSubCommand) *Message {
	var msg = NewMessage(constants.MESSAGE_SUB_COMMAND)
	msg.WriteIByte(int(command))
	return msg
}

func MessageNotLogin(command constants.MessageNotLogin) *Message {
	var msg = NewMessage(constants.NOT_LOGIN)
	msg.WriteIByte(int(command))
	return msg
}

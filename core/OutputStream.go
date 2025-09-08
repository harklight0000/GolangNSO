package core

import (
	"bytes"
	"go.uber.org/zap"
	"io"
	. "nso/logging"
	"nso/utils"
)

type OutputStream interface {
	WriteInt(int)
	WriteShort(int16)
	WriteIShort(int)
	WriteLong(int64)
	WriteUTF(string)
	WriteString(string)
	WriteFull([]byte)
	WriteIByte(int)
	WriteByte(byte)
	WriteSByte(int8)
	WriteBool(bool)
}

type DataOutputStream struct {
	io.Writer
}

func (d *DataOutputStream) WriteIShort(value int) {
	d.WriteShort(int16(value))
}

func (d *DataOutputStream) WriteByte(b byte) {
	_, err := d.Write([]byte{b})
	if err != nil {
		Logger.Panic("Error writing byte", zap.Error(err))
	}

}

func (d *DataOutputStream) WriteSByte(b int8) {
	_, err := d.Write([]byte{utils.Byte(b)})
	if err != nil {
		Logger.Panic("Error writing byte", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteBool(b bool) {
	if b {
		d.WriteIByte(1)
	} else {
		d.WriteIByte(0)
	}
}

func (d *DataOutputStream) WriteInt(value int) {
	_, err := d.Write([]byte{byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)})
	if err != nil {
		Logger.Panic("Error writing int", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteShort(value int16) {
	_, err := d.Write([]byte{byte(value >> 8), byte(value)})
	if err != nil {
		Logger.Panic("Error writing short", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteLong(value int64) {
	_, err := d.Write([]byte{byte(value >> 56), byte(value >> 48), byte(value >> 40), byte(value >> 32), byte(value >> 24), byte(value >> 16), byte(value >> 8), byte(value)})
	if err != nil {
		Logger.Panic("Error writing long", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteUTF(value string) {
	var buffer = bytes.Buffer{}
	d.WriteShort(int16(len(value)))
	buffer.WriteString(value)
	_, err := d.Write(buffer.Bytes())
	if err != nil {
		Logger.Panic("Error writing UTF", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteString(value string) {
	var buffer = bytes.Buffer{}
	d.WriteShort(int16(len(value)))
	buffer.WriteString(value)
	_, err := d.Write(buffer.Bytes())
	if err != nil {
		Logger.Panic("Error writing UTF", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteFull(value []byte) {
	_, err := d.Write(value)
	if err != nil {
		Logger.Panic("Error writing full", zap.Error(err))
	}
}

func (d *DataOutputStream) WriteIByte(value int) {
	_, err := d.Write([]byte{byte(value)})
	if err != nil {
		Logger.Panic("Error writing byte", zap.Error(err))
	}
}

func NewDataOutputStream(writer io.Writer) *DataOutputStream {
	return &DataOutputStream{Writer: writer}
}

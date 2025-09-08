package core

import (
	"bytes"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
	"io"
	. "nso/logging"
	"nso/utils"
	"strconv"
)

type InputStream interface {
	ReadInt() int
	ReadShort() int16
	ReadLong() int64
	ReadUTF() string
	ReadString() string
	ReadByte() int8
	ReadUByte() byte
	ReadUnsignedShort() uint16
	ReadUnsignedInt() uint32
	Available() int
	Skip(int) error
	ReadBoolean() bool
}

type DataInputStream struct {
	Reader io.Reader
}

func (d *DataInputStream) ReadBoolean() bool {
	return d.ReadByte() != 0
}

func (d *DataInputStream) Skip(i int) error {
	_, err := d.Reader.Read(make([]byte, i))
	if err != nil {
		return eris.Wrap(err, "Error skipping "+strconv.Itoa(i)+" bytes")
	}
	return nil
}

func (d *DataInputStream) Read(p []byte) (n int, err error) {
	read, err := d.Reader.Read(p)
	if err != nil {
		return -1, eris.Wrap(err, "Error reading bytes")
	}
	return read, nil
}

func NewDataInputStream(reader io.Reader) *DataInputStream {
	return &DataInputStream{Reader: reader}
}

func NewDataInputFromByte(buff []byte) *DataInputStream {
	var buffer bytes.Buffer
	buffer.Write(buff)
	return &DataInputStream{Reader: &buffer}
}

func (d *DataInputStream) Available() int {
	switch d.Reader.(type) {
	case *bytes.Buffer:
		return d.Reader.(*bytes.Buffer).Len()
	default:
		return 0
	}
}

func (d *DataInputStream) ReadUnsignedInt() uint32 {
	temp := make([]byte, 4)
	_, err := d.Reader.Read(temp)
	if err != nil {
		err := eris.Wrap(err, "Error when reading unsigned int")
		Logger.Panic("Error reading int ", zap.Error(err))
	}
	var result uint32
	result = uint32(temp[0]) << 24
	result |= uint32(temp[1]) << 16
	result |= uint32(temp[2]) << 8
	result |= uint32(temp[3]) << 0
	return result
}

func (d *DataInputStream) ReadInt() int {
	temp := make([]byte, 4)
	_, err := d.Reader.Read(temp)
	if err != nil {
		Logger.Panic("Error reading int", zap.Error(err))
	}
	var result int
	result = int(temp[0]) << 24
	result |= int(temp[1]) << 16
	result |= int(temp[2]) << 8
	result |= int(temp[3]) << 0
	return result
}

func (d *DataInputStream) ReadShort() int16 {
	tmp := make([]byte, 2)
	_, err := d.Reader.Read(tmp)
	if err != nil {
		Logger.Panic("Error reading short", zap.Error(err))
	}
	return int16(tmp[0])<<8 | int16(tmp[1])
}

func (d *DataInputStream) ReadLong() int64 {
	b := make([]byte, 8)
	_, err := d.Reader.Read(b)
	if err != nil {
		Logger.Panic("Error reading long", zap.Error(err))
	}
	return int64(b[0])<<56 | int64(b[1])<<48 | int64(b[2])<<40 | int64(b[3])<<32 | int64(b[4])<<24 | int64(b[5])<<16 | int64(b[6])<<8 | int64(b[7])
}

func (d *DataInputStream) ReadUTF() string {
	length := d.ReadUnsignedShort()
	tmp := make([]byte, length)
	_, err := d.Reader.Read(tmp)
	if err != nil {
		Logger.Panic("Error reading UTF", zap.Error(err))
	}
	return string(tmp)
}

func (d *DataInputStream) ReadString() string {
	return d.ReadUTF()
}

func (d *DataInputStream) ReadByte() int8 {
	tmp := make([]byte, 1)
	_, err := d.Reader.Read(tmp)
	if err != nil {
		panic(eris.Wrap(err, "Error reading byte"))
	}
	return int8(tmp[0])
}

func (d *DataInputStream) ReadUByte() byte {
	return utils.Byte(d.ReadByte())
}

func (d *DataInputStream) ReadUnsignedShort() uint16 {
	s := d.ReadShort()
	if s > 0 {
		return uint16(s)
	}
	return uint16(int(s) + 65536)
}

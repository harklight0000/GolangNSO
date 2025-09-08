package utils

func Byte(sByte int8) byte {
	if sByte > 0 {
		return byte(sByte)
	}
	r := int(sByte) + 256
	return byte(r & 0xFF)
}

func Short(sByte int8) byte {
	if sByte > 0 {
		return byte(sByte)
	}
	r := int(sByte) + 0xFFFF
	return byte(r & 0xFFFF)
}

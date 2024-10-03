package goutils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func BytesToInt(b []byte, isSymbol bool, bigEndian bool) (int, error) {
	if isSymbol {
		return bytesToIntS(b, bigEndian)
	}
	return bytesToIntU(b, bigEndian)
}
func BytesMustToInt(b []byte, isSymbol bool, bigEndian bool) int {
	if isSymbol {
		tmp, _ := bytesToIntS(b, bigEndian)
		return tmp
	}
	tmp, _ := bytesToIntU(b, bigEndian)
	return tmp
}

// 字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte, bigEndian bool) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	var err error
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return int(tmp), err
	case 2:
		var tmp uint16
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return int(tmp), err
	case 4:
		var tmp uint32
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// 字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte, bigEndian bool) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	var err error
	switch len(b) {
	case 1:
		var tmp int8
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
		}

		return int(tmp), err
	case 2:
		var tmp int16
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return int(tmp), err
	case 4:
		var tmp int32
		if bigEndian {
			err = binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			err = binary.Read(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// 整形转换成字节
func IntToBytes(n int64, b byte, bigEndian bool) ([]byte, error) {
	//n := m.(int64)
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		if bigEndian {
			binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		if bigEndian {
			binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		if bigEndian {
			binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return bytesBuffer.Bytes(), nil
	case 5, 6, 7, 8:
		tmp := int64(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		if bigEndian {
			binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		} else {
			binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
		}
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

func HexMust2Bytes(src string) []byte {
	b, _ := hex.DecodeString(src)
	return b
}

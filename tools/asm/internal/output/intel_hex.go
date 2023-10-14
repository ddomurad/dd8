package output

import (
	"encoding/hex"
	"strings"

	"github.com/ddomurad/dd8/tools/asm/internal"
)

func getHEXLine(addr int, dataline []byte) string {
	linebytes := []byte{
		byte(len(dataline)),
		byte(addr >> 8), byte(addr),
		0x00,
	}

	linebytes = append(linebytes, dataline...)
	cs := uint8(0)
	for _, b := range linebytes {
		cs += b
	}

	cs = uint8(uint16(0x100) - uint16(cs))
	linebytes = append(linebytes, cs)

	return ":" + strings.ToUpper(hex.EncodeToString(linebytes)) + "\n"
}

func GetIntelHEX(byteCode internal.ByteCode) []byte {
	output := ""
	dataLine := []byte{}
	addrs := byteCode.GetAddresses()
	lastAddr := addrs[0] - 1
	startAddr := addrs[0]

	for _, addr := range addrs {
		if lastAddr+1 != addr || len(dataLine) >= 0xff {
			output += getHEXLine(startAddr, dataLine)
			startAddr = addr
			dataLine = []byte{}
		}

		dataLine = append(dataLine, byteCode[addr])
		lastAddr = addr
	}

	output += getHEXLine(startAddr, dataLine)
	output += ":00000001FF\n"
	return []byte(output)
}

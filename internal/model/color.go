package model

import "errors"

// HexToRGB takes a hex color value and returns three ints with its red, green, and blue values.
// Source: https://stackoverflow.com/a/54200713/7687024
func HexToRGB(s string) (r, g, b int, err error) {
	var errInvalidFormat = errors.New("invalid format")

	if len(s) == 0 {
		return r, g, b, errInvalidFormat
	}
	if s[0] != '#' {
		return r, g, b, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	var rVal, gVal, bVal uint8
	switch len(s) {
	case 7:
		rVal = hexToByte(s[1])<<4 + hexToByte(s[2])
		gVal = hexToByte(s[3])<<4 + hexToByte(s[4])
		bVal = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		rVal = hexToByte(s[1]) * 17
		gVal = hexToByte(s[2]) * 17
		bVal = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}

	r = int(rVal)
	g = int(gVal)
	b = int(bVal)
	return
}

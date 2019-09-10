package lib

import "errors"

// HexToRGB takes a hex color value and returns three ints with its red, green, and blue values.
// Source: https://stackoverflow.com/a/54200713/7687024
func HexToRGB(s string) (r, g, b uint8, err error) {
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

	switch len(s) {
	case 7:
		r = hexToByte(s[1])<<4 + hexToByte(s[2])
		g = hexToByte(s[3])<<4 + hexToByte(s[4])
		b = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		r = hexToByte(s[1]) * 17
		g = hexToByte(s[2]) * 17
		b = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}

	return
}

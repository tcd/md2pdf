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

// DefaultFG rgb values.
func DefaultFG() (r, g, b int) {
	r = 36
	g = 41
	b = 46
	return
}

// DefaultBG rgb values.
func DefaultBG() (r, g, b int) {
	r = 255
	g = 255
	b = 255
	return
}

// LinkFG rgb values. (<a>)
func LinkFG() (r, g, b int) {
	r = 3
	g = 102
	b = 214
	return
}

// CodeSpanBG rgb values. (<code>)
func CodeSpanBG() (r, g, b int) {
	r = 243
	g = 243
	b = 243
	return
}

// CodeBlockBG rgb values. (<code>)
func CodeBlockBG() (r, g, b int) {
	r = 246
	g = 248
	b = 250
	return
}

// TableCellBG rgb values.
func TableCellBG() (r, g, b int) {
	r = 246
	g = 248
	b = 250
	return
}

// BlockquoteFG rgb values.
func BlockquoteFG() (r, g, b int) {
	r = 106
	g = 115
	b = 125
	return
}

// BlockquoteBorder rgb values.
func BlockquoteBorder() (r, g, b int) {
	r = 223
	g = 226
	b = 229
	return
}

// BlockquoteBG rgb values.
func BlockquoteBG() (r, g, b int) {
	r = 223
	g = 226
	b = 229
	return
}

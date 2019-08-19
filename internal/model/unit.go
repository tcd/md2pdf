package model

/**
 * These uints are all assuming 72 dpi.
 *
 * References:
 * https://graphicdesign.stackexchange.com/questions/199/point-vs-pixel-what-is-the-difference
 *
 * Alternatives:
 * https://godoc.org/golang.org/x/exp/shiny/unit
 * https://github.com/martinlindhe/unit
 */

const dpi = 72

// ============================================================================
// Inches
// ============================================================================

// IN represents a length in Inches.
type IN float64

// ToF returns a inch value as a float64.
func (in IN) ToF() float64 {
	return float64(in)
}

// ToMM converts Inches to Millimeters.
func (in IN) ToMM() MM {
	return MM(in * 25.4)
}

// ToPT converts Inches to Points.
func (in IN) ToPT() PT {
	return PT(in * dpi)
}

// ============================================================================
// Millimeters
// ============================================================================

// MM represents a length in Millimeters.
type MM float64

// ToF returns a millimeter value as a float64.
func (mm MM) ToF() float64 {
	return float64(mm)
}

// ToIN converts Millimeters to Inches.
func (mm MM) ToIN() IN {
	return IN(mm / 25.4)
}

// ToPT converts Millimeters to Points.
func (mm MM) ToPT() PT {
	return PT(mm / (25.4 / dpi))
}

// ============================================================================
// Points
// ============================================================================

// PT represents a length in Points.
type PT float64

// ToF returns a point value as a float64.
func (pt PT) ToF() float64 {
	return float64(pt)
}

// ToIN converts Points to Inches.
func (pt PT) ToIN() IN {
	return IN(pt / dpi)
}

// ToMM converts Points to Millimeters.
func (pt PT) ToMM() MM {
	return MM(pt * (25.4 / dpi))
}

// ============================================================================
// Misc
// ============================================================================

// PxToMm converts pixels to millimeters.
func PxToMm(pixels int) float64 {
	px := float64(pixels)
	return (px * (25.4 / dpi))
}

package types

// Line is a single printable line. It implements the printable interface.
type Line struct {
	Line      string
	LineColor string
}

// String returns the content of a line
func (l Line) String() string {
	return l.Line
}

// Color returns the color associated with the line
func (l Line) Color() string {
	return l.LineColor
}

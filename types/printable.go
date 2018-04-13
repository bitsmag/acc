package types

// Printable provides a printable string and a associated color
type Printable interface {
	String() string
	Color() string
}

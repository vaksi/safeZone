package services

// MultiplyService is struct service multiply
type MultiplyService struct {
	X int64
	Y int64
}

// Multiply is return from variable
func (m *MultiplyService) Multiply() (multiply int64){
	return m.X * m.Y
}

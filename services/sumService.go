package services

// SumService is struct service sum
type SumService struct {
	X int64
	Y int64
}

// Sum is return from variable
func (s *SumService) Sum() (sum int64){
	return s.X + s.Y
}
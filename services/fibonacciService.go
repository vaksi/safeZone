package services

type FibonacciService struct {
	N int64
}

// calculate of function generate fibonacci
func calculate() func() int64 {
	a, b := int64(0), int64(1)
	return func() int64 {
		a, b = b, a+b
		return a
	}
}

func (fs *FibonacciService) Fibonacci() (fib []int64) {
	f := calculate()
	for i := int64(0); i < fs.N; i++ {
		switch {
		case i == 0:
			fib = append(fib, i)
		default:
			fib = append(fib, f())
		}
	}
	return
}
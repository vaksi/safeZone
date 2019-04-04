package services

// PrimeService of service
type PrimeService struct {
	N int64
}

func generate(ch chan<- int64) {
	for i := int64(2); ; i++ {
		ch <- i
	}
}


func isPrime(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}


// Prime this service for get prime number
func (ps *PrimeService) Prime() (prime []int64) {
	// create channel
	ch := make(chan int64)

	// launch goroutine Generate
	go generate(ch)

	// selection prime
	for i := int64(0); i < ps.N; i++ {
		p := <-ch
		prime = append(prime, p)
		ch1 := make(chan int64)
		go isPrime(ch, ch1, p)
		ch = ch1
	}
	return
}
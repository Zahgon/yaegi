package main

func Generate(ch chan<- int) { _ = "STUB: not implemented"; return }

func Filter(in <-chan int, out chan<- int, prime int) { _ = "STUB: not implemented"; return }

func main() {
	ch := make(chan int)
	go Generate(ch)

	for i := 0; i < 10; i++ {
		prime := <-ch
		println(prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}

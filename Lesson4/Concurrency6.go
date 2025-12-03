// Timeout
func main() {
  var chan1 = make(chan int, 2)
  var timeout = time.After(5 * time.Second)
  // ...
  select {
    case v1 := <-chan1:
      fmt.Printf("Received %v from channel 1\n", v1)
    case <-timeout::
      fmt.Printf("Timeout\n")
  }
}

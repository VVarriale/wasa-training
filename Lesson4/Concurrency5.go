//Select
func main() {
  var chan1 = make(chan int, 2)
  var chan2 = make(chan int, 2)
  var chan3 = make(chan int, 2)
  // ...
  select {
    case v1 := <-chan1:
      fmt.Printf("Received %v from channel 1\n", v1)
    case v2 := <-chan2:
      fmt.Printf("Received %v from channel 2\n", v1)
    case chan3 <- 1:
      fmt.Printf("Sent value to channel 3\n")
    default:
      fmt.Printf("No one is ready to communicate\n", v1)
  }
}

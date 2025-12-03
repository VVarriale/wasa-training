// Buffered Channels
func main() {
  var channel = make(chan int, 2)
  go func() {
    var i = 0
    for i < 10 {
      channel <- i
      i++
    }
  }()
  var j = 0
  for j < 10 {
    fmt.Println(<-channel)
    j++
  }
}

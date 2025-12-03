func main() {
  go func() {
    var i = 0
    for i < 10 {
      fmt.Println(i)
      i++
    }
  }()
  var j = 0
  for j < 10 {
    fmt.Println(j)
    j++
  }
}

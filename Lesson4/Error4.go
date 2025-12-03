func main() {
  var someString = "12345"
  value, err := strconv.ParseInt(someString, 10, 64)
  if err != nil {
    panic(err)
  }
}

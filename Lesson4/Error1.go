// Errors have priority!
// Always handle errors!! Handle them immediately!
var someString string
// ...
value, err := strconv.ParseInt(someString, 10, 64)
if err != nil {
// What now?!?
}

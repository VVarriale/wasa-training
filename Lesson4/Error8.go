type MyError struct {
  LineNumber uint
  Message string
}

func (e *MyError) Error() string {
  return fmt.Sprintf("Error in line %d: %s", e.LineNumber, e.Message)}

func myFunction(str string) error {
  // ...
  return &MyError{LineNumber: 2, Message: "Missing :"}
}

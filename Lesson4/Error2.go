// Strategy 1: handle internally
// parseOrZero returns the parsed number,
// or zero if parse is not possible
func parseOrZero(someString string) int64 {
  value, err := strconv.ParseInt(someString, 10, 64)
  if err != nil {
    return 0
  }
  return value
}  

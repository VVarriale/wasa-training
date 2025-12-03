// Errors in function? Return to caller
// Strategy 2: pass the error to the caller
func parseAndIncrement(strValue string) (int64, error) {
  value, err := strconv.ParseInt(strValue, 10, 64)
  if err != nil {
    // Cleanup here if necessary, or use defer
    return 0, fmt.Errorf("error parsing value: %w", err)
  }
  value++
  return value, nil
}

var ErrOutOfRange = errors.New("value out of range")
// ParseLatitude parses the string as latitude and
// checks the data validity (latitude range)
func ParseLatitude(str string) (Latitude, error) {
  latitude, err := strconv.ParseFloat(str, 64)
  if err != nil {
    return 0, err
  } else if latitude < -90 || latitude > 90 {
    return 0, ErrOutOfRange
  }
  return Latitude(latitude), nil
}
// ...continue

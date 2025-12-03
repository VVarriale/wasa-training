// ...see previous slide
func main() {
  // ...
  latitude, err := ParseLatitude(latitudeInString)
  if errors.Is(err, ErrOutOfRange) {
    // Handle as invalid range
  } else {
  // Handle in another way
  }
}

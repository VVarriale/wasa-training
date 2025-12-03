type Latitude float64
func (lat *Latitude) IsValid() bool {
return lat != nil && -90 <= *lat && *lat <= 90
}
type Longitude float64
func (lng *Longitude) IsValid() bool {
return lng != nil && -180 <= *lng && *lng <= 180
}

type Vertex struct {
X int
Y int
}
func (v *Vertex) Equal(other *Vertex) bool {
return v.X == other.X && v.Y == other.Y
}
// Somewhere in the code
vtx := &Vertex{1, 2}
if vtx.Equal(&Vertex{2, 3}) {
// ...

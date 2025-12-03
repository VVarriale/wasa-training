//Struct implementing Stringer
type Vertex struct {
X int
Y int
}
func (v *Vertex) String() string {
return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}
// ...
func printSomething(s fmt.Stringer) {
fmt.Println(s.String())
}

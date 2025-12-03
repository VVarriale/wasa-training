//Writer
// io/io.go:96
type Writer interface {
Write(p []byte) (n int, err error)
}
// net/http/server.go:95
type ResponseWriter interface {
Header() Header
Write([]byte) (int, error)
WriteHeader(statusCode int)
}
func saveFile(w io.Writer, content []byte) {
w.Write(content)
}

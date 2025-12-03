// sync.Mutex
var mu sync.Mutex
var idx int

func Increment() {
  mu.Lock()
  defer mu.Unlock()
  idx++
}

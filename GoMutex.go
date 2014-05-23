  type contador struct {
    mu sync.Mutex
    x  int
  }
  func (c *contador) incrementar() (result int){
    c.mu.Lock() // HL
    defer c.mu.Unlock() // HL
    c.x += 1
    return c.x
  }


# Parallel

A super simple wrapper for using your synchronous code in a parallel fashion.




### Example


```go
func downloadChunk(chunk int) error {
  // do some downloading
  return nil
}

  func downloadFullFile(chunks []int) error {
    m := Manager{
      wg:   &sync.WaitGroup{},
      errs: []error{},
    }

  for _, i := range chunks{
    go m.Start(func() error {
      return tester(2)
    })
  }

  err := m.Return()
  if err != nil {
    t.Fatalf("non-nil error found from safe function")
  }
}
```



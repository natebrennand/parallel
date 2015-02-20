
# Parallel

[![Build Status](https://travis-ci.org/natebrennand/parallel.svg?branch=master)](https://travis-ci.org/natebrennand/parallel)

A super simple wrapper for using your synchronous code in a parallel fashion.



### Example


```go
func downloadChunk(chunk int) error {
  // do some downloading
  return nil
}

func downloadFullFile(chunks []int) error {
  m := parallel.DefaultManager()

  for _, i := range chunks{
    m.Start(func() error {
      return downloadChunk(i)
    })
  }

  // blocks until all calls are returned
  return m.Return()
}
```




# Parallel

[![GoDoc](https://godoc.org/github.com/natebrennand/parallel?status.svg)](https://godoc.org/github.com/natebrennand/parallel)
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

## Custom Error Aggregation

You can also define your own function to aggregate the list of errors into a single error.
By fulfilling the [Aggregator](http://godoc.org/github.com/natebrennand/parallel#Aggregator) signature you can create your own client with `CustomClient()`.


```go
fn = func(errs []error) error {
  strs := make([]string, len(errs))
  for i, e := range errs {
    strs[i] = e.Error()
  }
  return errors.New(strings.Join(strs, " + "))
}

m := parallel.CustomClient(fn)
```




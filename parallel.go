package parallel

import (
	"fmt"
	"strings"
	"sync"
)

// Manager manages multiple asynchronous function executions and aggregating any errors caused.
type Manager struct {
	wg   *sync.WaitGroup
	errs []error
	lock *sync.Mutex
	agg  Aggregator
}

// Aggregator is the signature of a function used to aggregate a possible list of errors
// into a single error to be digested by the client.
type Aggregator func([]error) error

func defaultAggregator(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	// accumulate all errors
	strs := make([]string, len(errs))
	for i, e := range errs {
		strs[i] = e.Error()
	}
	return fmt.Errorf("agg err: {%s}", strings.Join(strs, " | "))

}

// DefaultManager creates a general manager. It fits all use cases.
func DefaultManager() Manager {
	return Manager{
		wg:   &sync.WaitGroup{},
		errs: []error{},
		lock: &sync.Mutex{},
		agg:  defaultAggregator,
	}
}

// CustomClient creates a manager with a provided accumulator fn.
func CustomClient(fn Aggregator) Manager {
	return Manager{
		wg:   &sync.WaitGroup{},
		errs: []error{},
		lock: &sync.Mutex{},
		agg:  fn,
	}
}

// Start takes a wrapped function and calls it asynchronously
func (m *Manager) Start(f func() error) {
	m.wg.Add(1)
	go func() {
		err := f()
		if err != nil {
			m.lock.Lock()
			m.errs = append(m.errs, err)
			m.lock.Unlock()
		}
		m.wg.Done()
	}()
}

// Return blocks and aggregates any errors
func (m *Manager) Return() error {
	// block while waiting for all fn calls to finish
	m.wg.Wait()

	// call accumulator to aggregate errors
	return m.agg(m.errs)
}

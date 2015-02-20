package parallel

import (
	"fmt"
	"sync"
)

// Manager manages multiple asynchronous function executions and aggregating any errors caused.
type Manager struct {
	wg   *sync.WaitGroup
	errs []error
}

// DefaultManager creates a general manager. It fits all use cases.
func DefaultManager() Manager {
	return Manager{
		wg:   &sync.WaitGroup{},
		errs: []error{},
	}
}

// Start takes a wrapped function and calls it asynchronously
func (m *Manager) Start(f func() error) {
	m.wg.Add(1)
	go func() {
		if e := f(); e != nil {
			m.errs = append(m.errs, e)
		}
		m.wg.Done()
	}()
}

// Return blocks and aggregates any errors
func (m *Manager) Return() error {
	// block while waiting
	m.wg.Wait()

	// accumulate all errors needed
	if len(m.errs) > 0 {
		errStr := ""
		for _, e := range m.errs {
			errStr += e.Error()
		}
		return fmt.Errorf("ERR: {%s}", errStr)
	}

	return nil
}

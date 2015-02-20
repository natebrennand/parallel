package parallel

import (
	"fmt"
	"sync"
)

type Manager struct {
	wg   *sync.WaitGroup
	errs []error
}

func (m *Manager) Start(f func() error) {
	m.wg.Add(1)
	go func() {
		if e := f(); e != nil {
			m.errs = append(m.errs, e)
		}
		m.wg.Done()
	}()
}

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

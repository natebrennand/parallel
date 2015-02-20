package parallel

import (
	"fmt"
	"sync"
	"testing"
)

func tester(i int) error {
	return nil
}

func TestGeneralSuccess(t *testing.T) {
	m := DefaultManager()

	m.Start(func() error {
		return tester(2)
	})

	err := m.Return()
	if err != nil {
		t.Fatalf("non-nil error found from safe function")
	}
}

func tester2(i int) error {
	return fmt.Errorf("error %d", i)
}

func TestGeneral(t *testing.T) {
	m := Manager{
		wg:   &sync.WaitGroup{},
		errs: []error{},
	}

	m.Start(func() error {
		return tester(2)
	})
	m.Start(func() error {
		return tester2(401)
	})
	m.Start(func() error {
		return tester2(404)
	})
	m.Start(func() error {
		return tester2(500)
	})

	err := m.Return()
	if err == nil {
		t.Fatalf("nil error found from unsafe function calls => {%s}", err)
	}
}

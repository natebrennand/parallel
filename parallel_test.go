package parallel_test

import (
	"fmt"
	"testing"

	"github.com/natebrennand/parallel"
)

func tester(i int) error {
	return nil
}

func TestGeneralSuccess(t *testing.T) {
	m := parallel.DefaultManager()

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
	m := parallel.DefaultManager()

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

func TestCustom(t *testing.T) {
	testError := "test123"

	fn := func(errs []error) error {
		return fmt.Errorf(testError)
	}

	m := parallel.CustomClient(fn)
	m.Start(func() error {
		return tester(2)
	})
	m.Start(func() error {
		return tester2(401)
	})

	err := m.Return()
	if err.Error() != testError {
		t.Fatalf("Custom error handling not functioning properly\n\texpected: %s\n\treturned %s",
			testError, err.Error())
	}
}

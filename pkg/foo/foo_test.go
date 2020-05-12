package foo

import (
	"testing"

	"github.com/PapayaJuiceEQ/test_import/pkg/db"
)

func TestFoo(t *testing.T) {
	db.IncI()
	i := db.GetI()
	t.Log("FOO:", i)
}

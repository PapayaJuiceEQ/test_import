package foo

import (
	"log"
	"testing"

	"github.com/PapayaJuiceEQ/test_import/pkg/db"
)

func TestFoo(t *testing.T) {
	db.IncI()
	i := db.GetI()
	log.Print("FOO:", i)
}

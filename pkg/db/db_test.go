package db

import (
	"testing"
)

func TestDB(t *testing.T) {
	IncI()
	i := GetI()
	t.Log("DB:", i)
}

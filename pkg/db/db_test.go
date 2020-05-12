package db

import (
	"log"
	"testing"
)

func TestDB(t *testing.T) {
	IncI()
	i := GetI()
	log.Print(i)
}

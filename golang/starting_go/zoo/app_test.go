package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	expect := "Zoo Application"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

/*
go test -v
=== RUN   TestAppName
--- PASS: TestAppName (0.00s)
PASS
ok
*/

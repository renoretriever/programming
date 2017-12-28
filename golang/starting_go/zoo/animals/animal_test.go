package animals

import (
	"testing"
)

func TestElephantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestMonkeyFeed(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestRabbitFeed(t *testing.T) {
	//expect := "Grass"
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

/* OK
go test -v ./animals
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
=== RUN   TestMonkeyFeed
--- PASS: TestMonkeyFeed (0.00s)
=== RUN   TestRabbitFeed
--- PASS: TestRabbitFeed (0.00s)
PASS
*/

/* NG
go test -v ./animals
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
=== RUN   TestMonkeyFeed
--- PASS: TestMonkeyFeed (0.00s)
=== RUN   TestRabbitFeed
--- FAIL: TestRabbitFeed (0.00s)
        animal_test.go:30: Grass != Carrot
FAIL
exit status 1
*/

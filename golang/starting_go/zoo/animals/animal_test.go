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

/*
go test -v ./animals
=== RUN   TestElephantFeed
--- PASS: TestElephantFeed (0.00s)
=== RUN   TestMonkeyFeed
--- PASS: TestMonkeyFeed (0.00s)
=== RUN   TestRabbitFeed
--- PASS: TestRabbitFeed (0.00s)
PASS
*/

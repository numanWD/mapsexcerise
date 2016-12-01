package main

import "testing"

func TestCharliesTeam(t *testing.T) {
	actual := charliesTeam()
	expectLen := 4

	if len(actual) != expectLen {
		t.Errorf(
			"got %d, expected %d", actual, expectLen,
		)
	}
}

func TestPeopleOnFloorFive(t *testing.T) {
	actual := peopleOnFloorFive()
	expectLen := 6

	if len(actual) != expectLen {
		t.Errorf(
			"got %d, expected %d", actual, expectLen,
		)
	}
}

func TestPeopleWorkingInProduct(t *testing.T) {
	actual := peopleWorkingInProduct()
	expectLen := 1

	if len(actual) != expectLen {
		t.Errorf(
			"got %d, expected %d", actual, expectLen,
		)
	}
}

func BenchmarkAverage(b *testing.B) {
	for n := 0; n < b.N; n++ {
		charliesTeam()
	}
}

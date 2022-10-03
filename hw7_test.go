package hw6

import (
	"testing"
)

func Test1(t *testing.T) {
	N := 3
	a := 1
	M := 4
	b := 2
	P := N*M
	got := solvePairCongRelPrime(N,a,M,b)
	want := 10 

	if (got - want % P) != 0  {
		t.Errorf("You say that %d solves x = %d mod %d and x = %d mod %d but it is not equal to %d mod %d", got, a, N, b, M, want, P)
	}
}

func Test2(t *testing.T) {
	N := 5
	a := 2
	M := 7
	b := 5
	P := N*M
	got := solvePairCongRelPrime(N,a,M,b)
	want := 12 

	if (got - want % P) != 0  {
		t.Errorf("You say that %d solves x = %d mod %d and x = %d mod %d but it is not equal to %d mod %d", got, a, N, b, M, want, P)
	}
}

func Test3(t *testing.T) {
	N := 19
	a := 17
	M := 23
	b := 13
	P := N*M
	got := solvePairCongRelPrime(N,a,M,b)
	want := 36 

	if (got - want % P) != 0  {
		t.Errorf("You say that %d solves x = %d mod %d and x = %d mod %d but it is not equal to %d mod %d", got, a, N, b, M, want, P)
	}
}

func Test4(t *testing.T) {
	N := 59
	a := 46
	M := 71
	b := 34
	P := N*M
	got := solvePairCongRelPrime(N,a,M,b)
	want := 105 

	if (got - want % P) != 0  {
		t.Errorf("You say that %d solves x = %d mod %d and x = %d mod %d but it is not equal to %d mod %d", got, a, N, b, M, want, P)
	}
}

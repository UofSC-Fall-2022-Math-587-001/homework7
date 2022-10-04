package hw7

import (
	"testing"
)

func Test1(t *testing.T) {
	N := 10
	a := 3
	got := MillerRabinTest(N,a)
	want := "Composite"
	
	if got != want {
		t.Errorf("Testing N=%d with a=%d you find %q when you should find %q\n", N, a, got, want)
	}
}

func Test2(t *testing.T) {
	N := 15
	a := 3
	got := MillerRabinTest(N,a)
	want := "Composite"
	
	if got != want {
		t.Errorf("Testing N=%d with a=%d you find %q when you should find %q\n", N, a, got, want)
	}
}

func Test3(t *testing.T) {
	N := 21
	a := 20
	got := MillerRabinTest(N,a)
	want := "Test Fails"
	
	if got != want {
		t.Errorf("Testing N=%d with a=%d you find %q when you should find %q\n", N, a, got, want)
	}
}

func Test4(t *testing.T) {
	N := 19
	a := 4
	got := MillerRabinTest(N,a)
	want := "Test Fails"
	
	if got != want {
		t.Errorf("Testing N=%d with a=%d you find %q when you should find %q\n", N, a, got, want)
	}
}

func Test5(t *testing.T) {
	N := 33
	a := 2
	got := MillerRabinTest(N,a)
	want := "Composite"
	
	if got != want {
		t.Errorf("Testing N=%d with a=%d you find %q when you should find %q\n", N, a, got, want)
	}
}

package math

import "testing"

func TestAdd(t *testing.T){

    got := Add(4, 6)
    want := 10

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
	
}

func TestSub(t *testing.T){

    got := Subtract(8, 6)
    want := 2

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
	
}

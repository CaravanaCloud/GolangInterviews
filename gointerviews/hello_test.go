package main

import "testing"

func TestHello(t *testing.T) {
	got := hello()
	if got != "hello world" {
		t.Errorf("hello() = %s; want 'hello world'", got)
	}
}

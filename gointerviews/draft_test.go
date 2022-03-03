package main

import "testing"

func TestDraft(t *testing.T) {
	got := draft()
	if got != "hello world" {
		t.Errorf("hello() = %s; want 'hello world'", got)
	}
}

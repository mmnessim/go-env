package env

import (
	"testing"
)

func TestNoPath(t *testing.T) {
	e, _ := New()
	if e.Get("TEST") != "PASS" {
		t.Fatalf("expected PASS got %s", e.Get("TEST"))
	}
	if e.Get("TEST_2") != "Pass 2" {
		t.Fatalf("expected Pass 2 got %s", e.Get("TEST_2"))
	}
	if e.Get("TEST_3") != "blorg" {
		t.Fatalf("expected blorg got %s", e.Get("TEST_3"))
	}
	expected := `/.\!@#$%^&*()_+<>?:\`
	if e.Get("TEST_4") != expected {
		t.Fatalf("expected %s got %s", expected, e.Get("TEST_4"))
	}
}

func TestPath(t *testing.T) {
	e, _ := New(".env.local")
	if e.Get("TEST") != "PASS" {
		t.Fatalf("expected PASS got %s", e.Get("TEST"))
	}
	if e.Get("BLORG") != "BLARK" {
		t.Fatalf("expected PASS got %s", e.Get("BLORG"))
	}
	if e.Get("HELLO") != "WORLD" {
		t.Fatalf("expected PASS got %s", e.Get("HELLO"))
	}
}

func TestDuplicateKeys(t *testing.T) {
	e, _ := New(".env.dup")
	if e.Get("TEST") != "PASS" {
		t.Fatalf("expected PASS got %s", e.Get("TEST"))
	}
	if e.Get("TEST2") != "PASS" {
		t.Fatalf("expected PASS got %s", e.Get("TEST2"))
	}
}

func TestKeyDoesNotExist(t *testing.T) {
	_, err := New("doesnotexist")
	if err == nil {
		t.Fatalf("expected err, got err == nil")
	}
}

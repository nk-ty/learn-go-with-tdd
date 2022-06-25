package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func (t testing.TB, got, want string) {
        t.Helper()
        if got !=want {
            t.Errorf("got %q want %q", got, want)
        }
    }
    t.Run("Saying hello to people", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })
    t.Run("Saying 'Hello, world' when an empty string is supplied", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, world"
        assertCorrectMessage(t, got, want)
    })
    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
}
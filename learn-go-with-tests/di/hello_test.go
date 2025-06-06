package main

import (
	"bytes"
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper()は、このメソッドがヘルパーであることをテストスイートに伝えるために必要です。
		// こうすることで、テストが失敗したときに報告される行番号は、テストヘルパーの中ではなく 呼び出された関数 の中を示します。
		t.Helper()
		if got != want {
			t.Errorf("got %q != want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func TestGreet(t *testing.T) {
	// bytesパッケージのbufferタイプは Writerインターフェースを実装しています。
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

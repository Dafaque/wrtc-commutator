package tests

import (
	"commutator/commands"
	"fmt"
	"testing"
)

func TestSendOffer(t *testing.T) {
	var err error = commands.SendOffer(nil, []byte("to[asdfasdf] with[disis an offer]"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendAnswer(t *testing.T) {
	var err error = commands.SendAnswer(nil, []byte("to[asdfasdf] with[disis an answer]"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSendOnline(t *testing.T) {
	var err error = commands.Online(nil, []byte("with[salouronili] $$$salamo aleyyqoomo"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestParseArg(t *testing.T) {
	var expected string = "bar"
	var src []byte = []byte(fmt.Sprintf("foo[%s]", expected))
	actual, err := commands.ForTestingOnly("foo", &src)

	if err != nil {
		t.Fatal(err)
	}

	if actual != expected {
		t.Fatalf("expected %s, got %s", expected, actual)
	}

	if len(src) != 0 {
		t.Fatalf("src size expected %d, got %d", 0, len(src))
	}
}

func BenchmarkExec(b *testing.B) {
	commands.Exec(nil, []byte("+with[salouronili] $$$salamo aleyyqoomo"))
	commands.Exec(nil, []byte(">to[asdfasdf] with[disis an offer]"))
	commands.Exec(nil, []byte("<to[asdfasf] with[disisan ansua]"))
}

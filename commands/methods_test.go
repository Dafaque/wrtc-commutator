package commands

import (
	"bytes"
	"commutator/connection"
	"commutator/errcodes"
	"fmt"
	"testing"
)

func TestSendOffer(t *testing.T) {
	var err errcodes.ErrorCode = SendOffer(&connection.Connection{}, []byte("to[asdfasdf] with[disis an offer]"))
	if err != errcodes.ERROR_CODE_NONE {
		t.Fatal(err)
	}
}

func TestSendAnswer(t *testing.T) {
	var err errcodes.ErrorCode = SendAnswer(&connection.Connection{}, []byte("to[asdfasdf] with[disis an answer]"))
	if err != errcodes.ERROR_CODE_NONE {
		t.Fatal(err)
	}
}

func TestSendOnline(t *testing.T) {
	var err errcodes.ErrorCode = Online(&connection.Connection{}, []byte("with[salouronili] $$$salamo aleyyqoomo"))
	if err != errcodes.ERROR_CODE_NONE {
		t.Fatal(err)
	}
}

func TestParseArg(t *testing.T) {
	var expected []byte = []byte("bar")
	var src []byte = []byte(fmt.Sprintf("to[%s]", expected))
	actual, err := parseArg(ARG_TO, &src)

	if err != errcodes.ERROR_CODE_NONE {
		t.Fatal(err)
	}

	if !bytes.EqualFold(actual, expected) {
		t.Fatalf("expected %s, got %s", expected, actual)
	}

	if len(src) != 0 {
		t.Fatalf("src size expected %d, got %d", 0, len(src))
	}
}

func BenchmarkExec(b *testing.B) {
	Exec(&connection.Connection{}, []byte("+with[salouronili] $$$salamo aleyyqoomo"))
	Exec(&connection.Connection{}, []byte(">to[asdfasdf] with[disis an offer]"))
	Exec(&connection.Connection{}, []byte("<to[asdfasf] with[disisan ansua]"))
}

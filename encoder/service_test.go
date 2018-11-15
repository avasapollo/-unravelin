package encoder

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	e HashEncoder
)

func TestMain(m *testing.M) {
	e = NewHashEncoder()
	os.Exit(m.Run())
}

func TestHashEncoder_Encode(t *testing.T) {
	res := e.Encode("http://localhost")
	t.Log(res)

	check, err := e.Decode(res)
	require.Nil(t, err)
	require.Equal(t, check, "http://localhost")
}

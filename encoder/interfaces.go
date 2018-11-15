package encoder

type HashEncoder interface {
	Encode(s string) string
	Decode(s string) (string, error)
}

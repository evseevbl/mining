package decoders

type name string

func (n name) Name() string {
	return string(n)
}

const value = "value"


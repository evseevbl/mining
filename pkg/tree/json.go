package tree

type any struct {
	data []byte
}

func (a *any) UnmarshalJSON(data []byte) error {
	a.data = data
	return nil
}

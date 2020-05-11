package decoders

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProtoDecoder_decode(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		fname string
		exp   interface{}
	}{
		{
			fname: "testdata/person.buff",
			exp:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.fname, func(t *testing.T) {
			d := AsProtoMessage()
			data, err := ioutil.ReadFile(tc.fname)
			a.NoError(err, "cannot read test file")
			d.decode(data)
			// a.NoError(err)
			// a.Equal(tc.exp, exp)
		})
	}
}

func TestProtoDecoder_Decode(t *testing.T) {
	a := assert.New(t)

	testCases := []struct {
		fname string
		exp   interface{}
	}{
		{
			fname: "testdata/persons.buff",
			exp:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.fname, func(t *testing.T) {
			d := AsProtoMessage()
			data, err := ioutil.ReadFile(tc.fname)
			a.NoError(err, "cannot read test file")
			_, err = d.Decode(data)
			a.NoError(err)
			// a.Equal(tc.exp, exp)
		})
	}
}

func TestWrapNumbers(t *testing.T) {

	testCases := []struct {
		name string
		in   []byte
		out  []byte
	}{
		{

			name: "simple",
			in:   []byte(`abc3xyz678end`),
			out:   []byte(`abc"3"xyz"678"end`),
		},
		{
			name: "foo",

			in: []byte(`3: {1: "asdfasdf888ffg"}`),
			out: []byte(`"3": {"1": "asdfasdf888ffg"}`),
		},
		{
			name: "foo",

			in: []byte(`{4: ""}`),
			out: []byte(`{"4": ""}`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_out := wrapNumbers(tc.in)
			in := string(tc.in)
			out := string(_out)
			exp := string(tc.out)
			fmt.Println(in)
			fmt.Println(out)
			assert.Equal(t, exp, out)
		})
	}
}

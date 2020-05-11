package decoders

import (
	"bytes"
	"fmt"
	"os/exec"
)

type protoDecoder struct {
	name
}

func AsProtoMessage() *protoDecoder {
	return &protoDecoder{
		name: "proto",
	}
}

type msg struct {
}

func (m *msg) Reset() {
	_ = "reset"
}

func (m *msg) ProtoMessage() {
}

func (d *protoDecoder) Decode(data []byte) (map[string][]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("protoc", "--decode_raw")
	cmd.Stdin = bytes.NewBuffer(data)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	b := out.Bytes()
	b = bytes.ReplaceAll(b, []byte(`{
`,
	), []byte(`{ `))
	b = bytes.ReplaceAll(b, []byte{10}, []byte(`, `))
	b = bytes.ReplaceAll(b, []byte(` {`), []byte(`: {`))
	b = bytes.ReplaceAll(b, []byte(`, }`), []byte(`}`))
	b = bytes.ReplaceAll(b, []byte(`\`), []byte(`\\`))
	b = wrapNumbers(b)
	// _ = b
	// s := string(b)

	fmt.Println(string(b))
	return nil, nil
}

func wrapNumbers(data []byte) []byte {
	buf := make([]byte, 0, len(data))
	// pos := 0
	i := 0
	for ; ; {
		if i >= len(data) {
			break
		}

		switch data[i] {
		case '"':
			start := i
			i++
			for ; data[i] != '"' && i < len(data); {
				i++
			}
			end := i

			if end > start {
				buf = append(buf, data[start:end+1]...)
			}
			// pos = end
		case '0', '1','2','3','4','5','6','7','8','9':
			start := i
			for ; '0' <= data[i] && data[i] <= '9' && i < len(data); {
				i++
			}
			end := i
			// buf = append(buf, data[pos:start]...)
			buf = append(buf, []byte(`"`)...)
			buf = append(buf, data[start:end]...)
			buf = append(buf, []byte(`"`)...)
			buf = append(buf, data[i])
		default:
			buf = append(buf, data[i])
			// pos = i
		}
		i++
	}
	return buf
}

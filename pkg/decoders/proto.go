package decoders

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protowire"
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
func (m *msg) String() string { return "" }

func (d *protoDecoder) decode(data []byte) {
	buf := proto.NewBuffer(data)

	n, t, val := protowire.ConsumeField(buf.Unread())
	switch t {
	case protowire.Fixed64Type:
		i, _ := buf.DecodeFixed64()
		fmt.Printf("i=%d", i)
	case protowire.BytesType:
		s, _ := buf.DecodeStringBytes()
		fmt.Printf("s=%s\n", s)
	case protowire.VarintType:
		v, _ := buf.DecodeVarint()
		fmt.Printf("i=%d", v)
	case protowire.Fixed32Type:
		v, _ := buf.DecodeFixed32()
		fmt.Printf("i=%d", v)
	default:
		fmt.Println(n, t, val)
	}
	d.decode(buf.Unread())
	// var wire uint64 for shift := uint(0); ; shift += 7 {
	// 	wire |= uint64(b&0x7F) << shift
	// 	if b < 0x80 {
	// 		break
	// 	}
	// }

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
	fmt.Println(string(b))
	return map[string][]byte{
		value: b,
	}, nil
}

// ToDo доделать
func decodeValues(data []byte) {
	decodeToStruct := func(s string) map[string][]string {
		for _, part := range strings.Split(s, ",") {
			split := strings.Split(part, ":")
			if len(split) == 2 {
				fmt.Printf("%s=%s\n", split[0], split[1])
			}
		}
		return nil
	}
	s := string(data)
	start := strings.Index(s, "{")
	end := strings.Index(s, "}")
	val := s[start+1 : end]
	decodeToStruct(val)
	fmt.Println(val)
}

// ToDo доделать до json или выкинуть
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
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
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

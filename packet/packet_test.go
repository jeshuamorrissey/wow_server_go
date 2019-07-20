package packet

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestPacketLoader(t *testing.T) {
	{
		expectedBuffer := "\x01TEST\x04test"

		cpa := ClientPacketA{
			A: 1,
			B: [4]byte{'T', 'E', 'S', 'T'},
			D: []byte("test"),
		}

		buffer := new(bytes.Buffer)
		err := cpa.Process(MakeWriter(buffer))
		assert.NilError(t, err)
		assert.Equal(t, buffer.String(), expectedBuffer)

		cpaRead := ClientPacketA{}
		buffer = bytes.NewBufferString(expectedBuffer)
		err = cpaRead.Process(MakeReader(buffer))
		assert.NilError(t, err)
		assert.DeepEqual(t, cpa, cpaRead)
	}
}

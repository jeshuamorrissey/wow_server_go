package common

import (
	"fmt"
	"io"
)

// PadBigIntBytes takes as input an array of bytes and a size and ensures that the
// byte array is at least nBytes in length. \x00 bytes will be added to the end
// until the desired length is reached.
func PadBigIntBytes(data []byte, nBytes int) []byte {
	if len(data) > nBytes {
		return data[:nBytes]
	}

	currSize := len(data)
	for i := 0; i < nBytes-currSize; i++ {
		data = append(data, '\x00')
	}

	return data
}

// ReadBytes will read a specified number of bytes from a given buffer. If not all
// of the data is read (or there was an error), an error will be returned.
func ReadBytes(buffer io.Reader, length int) ([]byte, error) {
	data := make([]byte, length)
	n, err := buffer.Read(data)
	if err != nil {
		return nil, fmt.Errorf("error while reading bytes: %v", err)
	}

	if n != length {
		return nil, fmt.Errorf("short read: wanted %v bytes, got %v", length, n)
	}

	return data, nil
}

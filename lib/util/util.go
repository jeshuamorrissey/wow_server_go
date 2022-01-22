package util

import (
	"fmt"
	"io"
)

// ReverseBytes takes as input a byte array and returns a reversed version
// of it.
func ReverseBytes(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}

	return data
}

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

	if length > 0 {
		n, err := buffer.Read(data)
		if err != nil {
			return nil, fmt.Errorf("error while reading bytes: %v", err)
		}

		if n != length {
			return nil, fmt.Errorf("short read: wanted %v bytes, got %v", length, n)
		}
	}

	return data, nil
}

// Clamp takes a min, max and value and will return the value unless it is out of bounds,
// in which case it will return the bound.
func Clamp(min, value, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}

// Package moviehash implements the Media Player Classic hashing algorithm.
// More info can be found at:
// http://trac.opensubtitles.org/projects/opensubtitles/wiki/HashSourceCodes
package moviehash

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
	"strconv"
)

// The hash generator uses the first and last 65536 bytes of the file.
const BLOCKSIZE = 1 << 16

// ComputeFileHash by reading from the given filename. This function wraps
// ComputeHash and converts the result to a string.
func ComputeFileHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h, err := ComputeHash(file)
	if err != nil {
		return "", err
	}

	strhash := strconv.FormatUint(h, 16)

	return strhash, nil
}

// ComputeHash computes the hash. If something goes wrong it returns an error.
func ComputeHash(r io.ReadSeeker) (uint64, error) {
	var hash uint64 = 0
	var err error

	// Get checksum for the first block
	if hash, err = checksum(hash, r); err != nil {
		return 0, err
	}

	// Set file pointer to the last block of the file
	if _, err := r.Seek(-BLOCKSIZE, os.SEEK_END); err != nil {
		return 0, nil
	}

	// Get checksum for the last block
	if hash, err = checksum(hash, r); err != nil {
		return 0, err
	}

	// Get file size
	size, err := r.Seek(0, os.SEEK_END)
	if err != nil {
		return 0, err
	}

	hash = uint64(size) + hash

	return hash, nil
}

// checksum calculates the checksum for a 64K block.
func checksum(hash uint64, r io.Reader) (uint64, error) {
	var b uint64
	var bufreader io.Reader
	bufreader = bufio.NewReader(r)
	for i := 0; i < BLOCKSIZE/8; i++ {
		err := binary.Read(bufreader, binary.LittleEndian, &b)
		if err != nil {
			return 0, err
		}

		hash += b
	}
	return hash, nil
}

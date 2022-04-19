package lab3

import (
	"bytes"
)

const symbolCount = 256

type entry struct {
	code  uint64
	bytes []byte
}

type dictionary struct {
	entries []entry
	last    uint64
}

func (d *dictionary) initialise() {
	*d = dictionary{
		entries: make([]entry, symbolCount),
		last:    symbolCount}

	for i := range d.entries {
		d.entries[i] = entry{code: uint64(i), bytes: []byte{byte(i)}}
	}
}

func (d *dictionary) newEntry(b []byte) {
	d.entries = append(d.entries, entry{code: d.last, bytes: make([]byte, 0)})
	d.entries[d.last].bytes = append(d.entries[d.last].bytes, b...)
	d.last++
}

func (d *dictionary) find(b []byte) int {
	for _, e := range d.entries {
		if bytes.Equal(b, e.bytes) {
			return int(e.code)
		}
	}

	return -1
}

package lab3

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
		entries: make([]entry, 0),
		last:    0}

	for i := 0; i < symbolCount; i++ {
		d.entries = append(d.entries, entry{code: uint64(i), bytes: []byte{byte(i)}})
	}

	d.last = symbolCount
}

func (d *dictionary) newEntry(bytes []byte) {
	d.entries = append(d.entries, entry{code: d.last, bytes: bytes})
	d.last++
}

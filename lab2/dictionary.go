package lab2

type Dictionary struct {
	symbols    map[byte]uint16
	totalCount uint16
}

func (d *Dictionary) Initialise() {
	*d = Dictionary{
		symbols:    make(map[byte]uint16),
		totalCount: 256}

	for i := 0; i < 256; i++ {
		d.symbols[byte(i)] = 1
	}
}

func (d *Dictionary) Rescale() {
	d.totalCount = 0

	for i := 0; i < 256; i++ {
		d.symbols[byte(i)] = (d.symbols[byte(i)] + 1) / 2
		d.totalCount += d.symbols[byte(i)]
	}
}

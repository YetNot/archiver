package shannonfano

import "archiver/lib/compression/vlc/table"

type Generator struct{}

func NewGenertor() Generator {
	return Generator{}
}

type encodingTable map[rune]code

type code struct {
	Char     rune
	Quantity int
	Bit      uint32
}

type charStat map[rune]int

func (g Generator) NewTable(text string) table.EncodingTable {
	stat := newCharStat(text)
}

func build(stat charStat) {

}

func newCharStat(text string) charStat {
	res := make(charStat)

	for _, ch := range text {
		res[ch]++
	}

	return res
}

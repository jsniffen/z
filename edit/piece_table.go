package edit

import "strings"

type Entry struct {
	add    bool
	start  int
	length int
}

type PieceTable struct {
	originalBuffer []byte
	addBuffer      []byte
	table          []Entry
}

func (pt *PieceTable) String() string {
	var sb strings.Builder
	for _, entry := range pt.table {
		var b []byte
		if entry.add {
			b = pt.addBuffer[entry.start : entry.start+entry.length]
		} else {
			b = pt.originalBuffer[entry.start : entry.start+entry.length]
		}
		sb.Write(b)
	}
	return sb.String()
}

func NewPieceTable(init string) *PieceTable {
	pt := &PieceTable{
		originalBuffer: []byte(init),
		addBuffer:      make([]byte, 0),
		table:          make([]Entry, 0),
	}

	pt.table = append(pt.table, Entry{
		add:    false,
		start:  0,
		length: len(init),
	})

	return pt
}

func (pt *PieceTable) Delete(pos int) {
	currentPos := 0
	for i, entry := range pt.table {
		if currentPos == pos {
			pt.table[i].length -= 1
			pt.table[i].start += 1
			if pt.table[i].length == 0 {
				pt.table = append(pt.table[:i], pt.table[i+1:]...)
			}
			return
		} else if pos == currentPos+entry.length-1 {
			pt.table[i].length -= 1
			if pt.table[i].length == 0 {
				pt.table = append(pt.table[:i], pt.table[i+1:]...)
			}
			return
		} else if pos > currentPos && pos < currentPos+entry.length {
			pt.table = append(pt.table, Entry{})

			for j := len(pt.table) - 1; j > i+1; j-- {
				pt.table[j] = pt.table[j-1]
			}

			prev := Entry{
				add: entry.add,
				start: entry.start,
				length: pos - currentPos,
			}

			next := Entry{
				add: entry.add,
				start: prev.start+prev.length+1,
				length: entry.length-prev.length-1,
			}

			pt.table[i], pt.table[i+1] = prev, next

			return
		}

		currentPos += entry.length
	}
}

func (pt *PieceTable) Insert(s string, pos int) {
	newEntry := Entry{
		add:    true,
		start:  len(pt.addBuffer),
		length: len(s),
	}

	pt.addBuffer = append(pt.addBuffer, []byte(s)...)

	currentPos := 0
	for i, entry := range pt.table {
		if currentPos == pos {
			pt.table = append(pt.table[:i+1], pt.table[i:]...)
			pt.table[i] = newEntry
			return
		} else if pos > currentPos && pos < currentPos+entry.length {
			newLength := pos - currentPos
			pt.table = append(pt.table, Entry{}, Entry{})

			for j := len(pt.table) - 1; j-2 > i; j -= 1 {
				pt.table[j] = pt.table[j-2]
			}

			pt.table[i+1] = newEntry
			pt.table[i+2] = Entry{
				add:    entry.add,
				start:  entry.start + newLength,
				length: entry.length - newLength,
			}

			pt.table[i].length = newLength
			return
		}

		currentPos += entry.length
	}
}

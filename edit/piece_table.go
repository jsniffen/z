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

func shiftRight(entries []Entry, start, n int) []Entry {
	for i := 0; i < n; i++ {
		entries = append(entries, Entry{})
	}

	for i := len(entries) - 1; i >= start+n; i-- {
		entries[i] = entries[i-n]
	}

	return entries
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

func (pt *PieceTable) Insert(b byte, pos int) {
	newEntry := Entry{
		add:    true,
		start:  len(pt.addBuffer),
		length: 1,
	}

	pt.addBuffer = append(pt.addBuffer, b)

	start := 0
	for i, entry := range pt.table {
		end := start + entry.length

		if pos == start {
			pt.table = shiftRight(pt.table, i, 1)
			pt.table[i] = newEntry
		} else if pos == end-1 {
			pt.table = shiftRight(pt.table, i, 1)
			pt.table[i+1] = newEntry
		} else if pos > start && pos < end {
			pt.table = shiftRight(pt.table, i, 2)

			prev := Entry{
				add:    entry.add,
				start:  entry.start,
				length: pos - start,
			}

			next := Entry{
				add:    entry.add,
				start:  prev.start + prev.length,
				length: entry.length - prev.length,
			}

			pt.table[i] = prev
			pt.table[i+1] = newEntry
			pt.table[i+2] = next
		}

		start += entry.length
	}
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
				add:    entry.add,
				start:  entry.start,
				length: pos - currentPos,
			}

			next := Entry{
				add:    entry.add,
				start:  prev.start + prev.length + 1,
				length: entry.length - prev.length - 1,
			}

			pt.table[i], pt.table[i+1] = prev, next

			return
		}

		currentPos += entry.length
	}
}

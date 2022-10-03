package edit

import "strings"

type Entry struct {
	add    bool
	start  int
	length int
}

func split(e Entry, i int) (Entry, Entry) {
	prev := Entry{
		add:    e.add,
		start:  e.start,
		length: i,
	}

	next := Entry{
		add:    e.add,
		start:  e.start + i,
		length: e.length - i,
	}

	return prev, next
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
			pt.table[i], pt.table[i+2] = split(entry, pos-start)
			pt.table[i+1] = newEntry
		}

		start += entry.length
	}
}

func (pt *PieceTable) Delete(pos int) {
	start := 0
	for i, entry := range pt.table {
		end := start + entry.length

		if pos == start {
			pt.table[i].start += 1
			pt.table[i].length -= 1
			if pt.table[i].length == 0 {
				pt.table = append(pt.table[:i], pt.table[i+1:]...)
			}
			break
		} else if pos == end-1 {
			pt.table[i].length -= 1
			if pt.table[i].length == 0 {
				pt.table = append(pt.table[:i], pt.table[i+1:]...)
			}
			break
		} else if pos > start && pos < end {
			pt.table = shiftRight(pt.table, i, 1)
			pt.table[i], pt.table[i+1] = split(entry, pos-start)
			pt.table[i+1].start += 1
			pt.table[i+1].length -= 1
			break
		}

		if pt.table[i].length == 0 {
			pt.table = append(pt.table[:i], pt.table[i+1:]...)
		}

		start += entry.length
	}
}

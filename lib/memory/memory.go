package memory

import (
	"io"
)

type Memory interface {
	Read(addr int) byte
	ReadWord(addr int) int
	Write(addr int, data byte)
	Length() int
}

type MemoryMapEntry struct {
	BaseAddr int
	LastAddr int
	Block    Memory
}

type MemoryMap struct {
	Entries []MemoryMapEntry
}

// This can be anything, right now it's this
func NewMemoryMap(rom io.Reader) *MemoryMap {
	entries := []MemoryMapEntry{
		{
			BaseAddr: 0,
			Block:    NewRandomAccessMemory(0x1000),
		},
		{
			BaseAddr: 0x1000,
			Block:    NewReadOnlyMemory(0x1000, rom),
		},
		{
			BaseAddr: 0x2000,
			Block:    NewRandomAccessMemory(65536 - 0x2000),
		},
	}

	for i, entry := range entries {
		entries[i].LastAddr = entry.BaseAddr + entry.Block.Length()
	}

	return &MemoryMap{
		Entries: entries,
	}
}

func (m *MemoryMap) resolve(addr int) (int, Memory) {
	for _, entry := range m.Entries {
		if addr < entry.LastAddr {
			return addr - entry.BaseAddr, entry.Block
		}
	}
	return addr, nil
}

func (m *MemoryMap) Read(addr int) byte {
	relAddr, block := m.resolve(addr)
	return block.Read(relAddr)
}

func (m *MemoryMap) ReadWord(addr int) int {
	relAddr, block := m.resolve(addr)
	return block.ReadWord(relAddr)
}

func (m *MemoryMap) Write(addr int, data byte) {
	relAddr, block := m.resolve(addr)
	block.Write(relAddr, data)
}

func (m *MemoryMap) Length() int {
	return 65536
}

func NewNesMemoryMap() *MemoryMap {
	entries := []MemoryMapEntry{
		{
			BaseAddr: 0,
			Block: NewMirrored(
				NewRandomAccessMemory(0x0800),
				0x7ff,
				0x2000,
			),
		},
		{
			BaseAddr: 0x2000,
			Block: NewMirrored(
				NewRandomAccessMemory(0x8),
				0x7,
				0x2000,
			),
		},
		{
			BaseAddr: 0x4000,
			Block:    NewRandomAccessMemory(65536 - 0x4000),
		},
	}

	for i, entry := range entries {
		entries[i].LastAddr = entry.BaseAddr + entry.Block.Length()
	}

	return &MemoryMap{
		Entries: entries,
	}
}

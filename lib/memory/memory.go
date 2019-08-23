package memory

import (
	"encoding/binary"
	"io"
)

type Memory interface {
	Read(addr int) byte
	ReadWord(addr int) int
	Write(addr int, data byte)
	Length() int
}

type RandomAccessMemory struct {
	Contents []byte
}

func NewRandomAccessMemory(size int) *RandomAccessMemory {
	return &RandomAccessMemory{
		Contents: make([]byte, size),
	}
}

func (r *RandomAccessMemory) Read(addr int) byte {
	return r.Contents[addr]
}

func (r *RandomAccessMemory) ReadWord(addr int) int {
	return int(binary.LittleEndian.Uint16(r.Contents[addr:]))
}

func (r *RandomAccessMemory) Write(addr int, data byte) {
	r.Contents[addr] = data
}

func (r *RandomAccessMemory) Length() int {
	return len(r.Contents)
}

type ReadOnlyMemory struct {
	Contents []byte
}

func NewReadOnlyMemory(size int, data io.Reader) *ReadOnlyMemory {
	contents := make([]byte, size)

	_, err := data.Read(contents)
	if err != nil {
		panic(err)
	}

	return &ReadOnlyMemory{
		Contents: contents,
	}
}

func (r *ReadOnlyMemory) Read(addr int) byte {
	return r.Contents[addr]
}

func (r *ReadOnlyMemory) ReadWord(addr int) int {
	return int(binary.LittleEndian.Uint16(r.Contents[addr:]))
}

func (r *ReadOnlyMemory) Write(addr int, data byte) {
	// we do nothing for now
}

func (r *ReadOnlyMemory) Length() int {
	return len(r.Contents)
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
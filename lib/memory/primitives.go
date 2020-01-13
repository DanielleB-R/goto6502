package memory

import (
	"encoding/binary"
	"io"
)

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

type MirroredMemory struct {
	underlying Memory
	mask       int
	size       int
}

func NewMirrored(underlying Memory, mask int, size int) *MirroredMemory {
	return &MirroredMemory{underlying, mask, size}
}

func (m *MirroredMemory) Read(addr int) byte {
	return m.underlying.Read(addr & m.mask)
}

func (m *MirroredMemory) ReadWord(addr int) int {
	return m.underlying.ReadWord(addr & m.mask)
}

func (m *MirroredMemory) Write(addr int, data byte) {
	m.underlying.Write(addr&m.mask, data)
}

func (m *MirroredMemory) Length() int {
	return m.size
}

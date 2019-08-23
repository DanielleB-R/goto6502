package memory

import "encoding/binary"

type Memory interface {
	Read(addr int) byte
	ReadWord(addr int) int
	Write(addr int, data byte)
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

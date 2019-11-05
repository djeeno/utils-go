package utils

type ByteOrder interface {
	Get16Bytes([]byte) [16]byte
	Put16Bytes([]byte, [16]byte)
	String() string
}

var Binary = binaryT{
	BigEndian:    bigEndianT{},
	LittleEndian: littleEndianT{},
}

type binaryT struct {
	BigEndian    bigEndianT
	LittleEndian littleEndianT
}

type bigEndianT struct{}

func (bigEndianT) String() string { return "BigEndian" }

func (bigEndianT) Get16Bytes(b []byte) [16]byte {
	_ = b[15] // bounds check hint to compiler; see golang.org/issue/14808
	return [16]byte{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7], b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15]}
}

func (bigEndianT) Put16Bytes(b []byte, v [16]byte) {
	_ = b[15] // early bounds check to guarantee safety of writes below
	b[0] = v[0]
	b[1] = v[1]
	b[2] = v[2]
	b[3] = v[3]
	b[4] = v[4]
	b[5] = v[5]
	b[6] = v[6]
	b[7] = v[7]
	b[8] = v[8]
	b[9] = v[9]
	b[10] = v[10]
	b[11] = v[11]
	b[12] = v[12]
	b[13] = v[13]
	b[14] = v[14]
	b[15] = v[15]
}

type littleEndianT struct{}

func (littleEndianT) String() string { return "LittleEndian" }

func (littleEndianT) Get16Bytes(b []byte) [16]byte {
	_ = b[15] // bounds check hint to compiler; see golang.org/issue/14808
	return [16]byte{b[15], b[14], b[13], b[12], b[11], b[10], b[9], b[8], b[7], b[6], b[5], b[4], b[3], b[2], b[1], b[0]}
}

func (littleEndianT) Put16Bytes(b []byte, v [16]byte) {
	_ = b[15] // early bounds check to guarantee safety of writes below
	b[15] = v[0]
	b[14] = v[1]
	b[13] = v[2]
	b[12] = v[3]
	b[11] = v[4]
	b[10] = v[5]
	b[9] = v[6]
	b[8] = v[7]
	b[7] = v[8]
	b[6] = v[9]
	b[5] = v[10]
	b[4] = v[11]
	b[3] = v[12]
	b[2] = v[13]
	b[1] = v[14]
	b[0] = v[15]
}

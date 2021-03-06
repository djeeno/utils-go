package utils

import (
	"bytes"
	"testing"
)

var (
	test16Bytes = [16]byte{0x03, 0x01, 0x04, 0x01, 0x05, 0x09, 0x02, 0x06, 0x05, 0x03, 0x05, 0x08, 0x09, 0x07, 0x09, 0x03}
)

func TestBinary(t *testing.T) {
	_ = Binary()
}

func TestBigEndian_String(t *testing.T) {
	t.Helper()
	testBinary := _binary

	if testBinary.BigEndian.String() != "BigEndian" {
		t.Errorf("TestBigEndian_String(): testBinary.BigEndian.String() failed.\n")
	}
}

func TestBigEndian_Get16Bytes(t *testing.T) {
	t.Helper()
	testBinary := _binary

	testBigEndianSliceGet := []byte{0x0f, 0x03, 0x01, 0x04, 0x01, 0x05, 0x09, 0x02, 0x06, 0x05, 0x03, 0x05, 0x08, 0x09, 0x07, 0x09, 0x03, 0x0f}

	if test16Bytes != testBinary.BigEndian.Get16Bytes(testBigEndianSliceGet[1:17]) {
		t.Errorf("TestBigEndian_Get16Bytes(): testBinary.BigEndian.Get16Bytes() failed.\n")
	}
}

func TestBigEndian_Put16Bytes(t *testing.T) {
	t.Helper()
	testBinary := _binary

	testBigEndianSlicePut := []byte{0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f}
	testSlicePutResult := []byte{0x0f, 0x03, 0x01, 0x04, 0x01, 0x05, 0x09, 0x02, 0x06, 0x05, 0x03, 0x05, 0x08, 0x09, 0x07, 0x09, 0x03, 0x0f}

	testBinary.BigEndian.Put16Bytes(testBigEndianSlicePut[0:16], test16Bytes)

	if bytes.Equal(testBigEndianSlicePut, testSlicePutResult) {
		t.Errorf("TestBigEndian_Put16Bytes(): testBinary.BigEndian.Put16Bytes() failed.\n")
	}
}

func TestLittleEndianT_String(t *testing.T) {
	t.Helper()
	testBinary := _binary

	if testBinary.LittleEndian.String() != "LittleEndian" {
		t.Errorf("TestLittleEndianT_String(): testBinary.LittleEndian.String() failed.\n")
	}
}

func TestLittleEndianT_Get16Bytes(t *testing.T) {
	t.Helper()
	testBinary := _binary

	testLittleEndianSliceGet := []byte{0x0f, 0x03, 0x09, 0x07, 0x09, 0x08, 0x05, 0x03, 0x05, 0x06, 0x02, 0x09, 0x05, 0x01, 0x04, 0x01, 0x03, 0x0f}

	if test16Bytes != testBinary.LittleEndian.Get16Bytes(testLittleEndianSliceGet[1:17]) {
		t.Errorf("TestLittleEndianT_Get16Bytes(): testBinary.LittleEndian.Get16Bytes() failed.\n")
	}
}

func TestLittleEndianT_Put16Bytes(t *testing.T) {
	t.Helper()
	testBinary := _binary

	testLittleEndianSlicePut := []byte{0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f, 0x0f}
	testLittleEndianSlicePutResult := []byte{0x0f, 0x03, 0x09, 0x07, 0x09, 0x08, 0x05, 0x03, 0x05, 0x06, 0x02, 0x09, 0x05, 0x01, 0x04, 0x01, 0x03, 0x0f}

	testBinary.LittleEndian.Put16Bytes(testLittleEndianSlicePut[0:16], test16Bytes)

	if bytes.Equal(testLittleEndianSlicePut, testLittleEndianSlicePutResult) {
		t.Errorf("TestLittleEndianT_Put16Bytes(): testBinary.LittleEndian.Put16Bytes() failed.\n")
	}
}

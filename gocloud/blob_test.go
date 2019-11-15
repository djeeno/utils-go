package gocloud

import (
	"context"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
	"io"
	"os"
	"testing"
)

// TODO: 何もかもを直す

func TestBlobT_PutFile(t *testing.T) {
	t.Helper()
	testFilePath := "/tmp/test.txt"
	testBlobPath := "tmp/test.txt"

	t.Run("normal/TODO", func(t *testing.T) {
		testBlob := Blob
		testBlob.blobWriterCloseFn = func(blobWriter *blob.Writer) error {
			return nil
		}
		testBlob.blobWriterWriteFn = func(blobWriter *blob.Writer, p []byte) (n int, err error) {
			return 0, nil
		}
		testBlob.bucketNewWriterFn = func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
			return nil, nil
		}
		testBlob.osOpenFn = func(name string) (file *os.File, e error) {
			return nil, nil
		}
		testBlob.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, nil
		}
		if err := testBlob.PutFile(context.TODO(), testFilePath, &blob.Bucket{}, testBlobPath, nil); err != nil {
			t.Errorf("TestBlobT_PutFile(): testBlob.PutFile(): err != nil: %v", err)
		}
	})

	t.Run("non-normal/blobWriter.Close", func(t *testing.T) {
		testBlob := Blob
		testBlob.blobWriterCloseFn = func(blobWriter *blob.Writer) error {
			return ErrorDummyErrorForTest
		}
		testBlob.blobWriterWriteFn = func(blobWriter *blob.Writer, p []byte) (n int, err error) {
			return 0, nil
		}
		testBlob.bucketNewWriterFn = func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
			return nil, nil
		}
		testBlob.osOpenFn = func(name string) (file *os.File, e error) {
			return nil, nil
		}
		testBlob.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, nil
		}
		if err := testBlob.PutFile(context.TODO(), testFilePath, &blob.Bucket{}, testBlobPath, nil); err == nil {
			t.Errorf("TestBlobT_PutFile(): testBlob.PutFile(): err == nil")
		}

	})

	t.Run("non-normal/blobWriter.Write", func(t *testing.T) {
		testBlob := Blob
		testBlob.blobWriterCloseFn = func(blobWriter *blob.Writer) error {
			return nil
		}
		testBlob.blobWriterWriteFn = func(blobWriter *blob.Writer, p []byte) (n int, err error) {
			return 0, ErrorDummyErrorForTest
		}
		testBlob.bucketNewWriterFn = func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
			return nil, nil
		}
		testBlob.osOpenFn = func(name string) (file *os.File, e error) {
			return nil, nil
		}
		testBlob.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, nil
		}
		if err := testBlob.PutFile(context.TODO(), testFilePath, &blob.Bucket{}, testBlobPath, nil); err == nil {
			t.Errorf("TestBlobT_PutFile(): testBlob.PutFile(): err == nil")
		}
	})

	t.Run("non-normal/ioutil.ReadAll", func(t *testing.T) {
		testBlob := Blob
		testBlob.blobWriterCloseFn = func(blobWriter *blob.Writer) error {
			return nil
		}
		testBlob.blobWriterWriteFn = func(blobWriter *blob.Writer, p []byte) (n int, err error) {
			return 0, nil
		}
		testBlob.bucketNewWriterFn = func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
			return nil, nil
		}
		testBlob.osOpenFn = func(name string) (file *os.File, e error) {
			return nil, nil
		}
		testBlob.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := testBlob.PutFile(context.TODO(), testFilePath, &blob.Bucket{}, testBlobPath, nil); err == nil {
			t.Errorf("TestBlobT_PutFile(): testBlob.PutFile(): err == nil")
		}
	})

	t.Run("non-normal/os.Open", func(t *testing.T) {
		testBlob := Blob
		testBlob.blobWriterCloseFn = func(blobWriter *blob.Writer) error {
			return nil
		}
		testBlob.blobWriterWriteFn = func(blobWriter *blob.Writer, p []byte) (n int, err error) {
			return 0, nil
		}
		testBlob.bucketNewWriterFn = func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
			return nil, nil
		}
		testBlob.osOpenFn = func(name string) (file *os.File, e error) {
			return nil, ErrorDummyErrorForTest
		}
		testBlob.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, nil
		}
		if err := testBlob.PutFile(context.TODO(), testFilePath, &blob.Bucket{}, testBlobPath, nil); err == nil {
			t.Errorf("TestBlobT_PutFile(): testBlob.PutFile(): err == nil")
		}
	})

	t.Run("non-normal/bucket.NewWriter", func(t *testing.T) {
		testBlob := Blob
		testBlob.blobWriterCloseFn = func(blobWriter *blob.Writer) error {
			return nil
		}
		testBlob.blobWriterWriteFn = func(blobWriter *blob.Writer, p []byte) (n int, err error) {
			return 0, nil
		}
		testBlob.bucketNewWriterFn = func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
			return nil, ErrorDummyErrorForTest
		}
		testBlob.osOpenFn = func(name string) (file *os.File, e error) {
			return nil, nil
		}
		testBlob.ioutilReadAllFn = func(r io.Reader) (bytes []byte, e error) {
			return nil, nil
		}
		if err := testBlob.PutFile(context.TODO(), testFilePath, &blob.Bucket{}, testBlobPath, nil); err == nil {
			t.Errorf("TestBlobT_PutFile(): testBlob.PutFile(): err == nil")
		}
	})

}

func TestBlobT_bucketNewWriter(t *testing.T) {
	t.Helper()
	testBlobPath := "tmp/test.txt"

	t.Run("non-normal/bucket.NewWriter", func(t *testing.T) {
		testBlob := Blob
		testContext := context.TODO()
		defer func() { _ = recover() }()
		bw, err := testBlob.bucketNewWriterFn(&blob.Bucket{}, testContext, testBlobPath, nil)
		defer func() { _ = bw.Close() }()
		if err != nil {
			t.Errorf("TestBlobT_bucketNewWriter(): testBlob.bucketNewWriterFn(): err != nil")
		}
	})
}

func TestBlobT_blobWriterWrite(t *testing.T) {
	t.Helper()

	t.Run("non-normal/bucket.NewWriter", func(t *testing.T) {
		testBlob := Blob
		defer func() { _ = recover() }()
		if _, err := testBlob.blobWriterWriteFn(&blob.Writer{}, []byte{}); err != nil {
			t.Errorf("TestBlobT_bucketNewWriter(): testBlob.blobWriterWriteFn(): err != nil")
		}
	})
}

func TestBlobT_blobWriterClose(t *testing.T) {
	t.Helper()

	t.Run("non-normal/bucket.NewWriter", func(t *testing.T) {
		testBlob := Blob
		defer func() { _ = recover() }()
		if err := testBlob.blobWriterCloseFn(&blob.Writer{}); err != nil {
			t.Errorf("TestBlobT_bucketNewWriter(): testBlob.bucketNewWriterFn(): err != nil")
		}
	})
}

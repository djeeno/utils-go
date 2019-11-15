package gocloud

import (
	"context"
	"gocloud.dev/blob"
	"io"
	"io/ioutil"
	"os"
)

var Blob = blobT{
	blobWriterCloseFn: blobWriterClose,
	blobWriterWriteFn: blobWriterWrite,
	bucketNewWriterFn: bucketNewWriter,
	ioutilReadAllFn:   ioutil.ReadAll,
	osOpenFn:          os.Open,
}

type blobT struct {
	blobWriterCloseFn func(blobWriter *blob.Writer) error
	blobWriterWriteFn func(blobWriter *blob.Writer, p []byte) (n int, err error)
	bucketNewWriterFn func(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error)
	ioutilReadAllFn   func(r io.Reader) ([]byte, error)
	osOpenFn          func(name string) (*os.File, error)
}

func bucketNewWriter(bucket *blob.Bucket, ctx context.Context, key string, opts *blob.WriterOptions) (_ *blob.Writer, err error) {
	return bucket.NewWriter(ctx, key, opts)
}

func blobWriterWrite(blobWriter *blob.Writer, p []byte) (n int, err error) {
	return blobWriter.Write(p)
}

func blobWriterClose(blobWriter *blob.Writer) error {
	return blobWriter.Close()
}

func (b *blobT) PutFile(ctx context.Context, filePath string, bucket *blob.Bucket, blobPath string, opts *blob.WriterOptions) error {
	blobWriter, err := b.bucketNewWriterFn(bucket, ctx, blobPath, opts)
	defer func() { _ = b.blobWriterCloseFn(blobWriter) }()
	if err != nil {
		return err
	}
	file, err := b.osOpenFn(filePath)
	if err != nil {
		return err
	}
	fileContent, err := b.ioutilReadAllFn(file)
	if err != nil {
		return err
	}
	if _, err := b.blobWriterWriteFn(blobWriter, fileContent); err != nil {
		return err
	}
	if err := b.blobWriterCloseFn(blobWriter); err != nil {
		return err
	}
	return nil
}

package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ZIP = zipT{
	filepathWalkFn:    filepath.Walk,
	ioCopyFn:          io.Copy,
	osCreateFn:        os.Create,
	osOpenFn:          os.Open,
	zipWriterCreateFn: zipWriterCreate,
}

type zipT struct {
	filepathWalkFn    func(root string, walkFn filepath.WalkFunc) error
	ioCopyFn          func(dst io.Writer, src io.Reader) (written int64, err error)
	osCreateFn        func(name string) (*os.File, error)
	osOpenFn          func(name string) (*os.File, error)
	zipWriterCreateFn func(w *zip.Writer, name string) (io.Writer, error)
}

func zipWriterCreate(w *zip.Writer, name string) (io.Writer, error) {
	return w.Create(name)
}

func (z *zipT) ArchivesRecursive(zipFilePath, directoryPathToArchive string, withoutRootDirectory bool) error {
	zipFile, err := z.osCreateFn(zipFilePath)
	if err != nil {
		return err
	}

	zw := zip.NewWriter(zipFile)

	walkFunc := func(filePath string, info os.FileInfo, err error) error {
		return walkFuncForArchivesRecursive(z, zw, filepath.Dir(directoryPathToArchive)+"/", filePath, info, err)
	}
	if withoutRootDirectory {
		walkFunc = func(filePath string, info os.FileInfo, err error) error {
			return walkFuncForArchivesRecursive(z, zw, directoryPathToArchive+"/", filePath, info, err)
		}
	}

	if err := z.filepathWalkFn(directoryPathToArchive, walkFunc); err != nil {
		return err
	}

	return zw.Close()
}

func walkFuncForArchivesRecursive(z *zipT, zw *zip.Writer, prefixTrimmingRelativePath string, filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	relativePath := strings.TrimPrefix(filePath, prefixTrimmingRelativePath)
	if info.IsDir() {
		relativePath = relativePath + "/" // directory
	}
	fileInZipWriter, err := z.zipWriterCreateFn(zw, relativePath)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}
	fsFile, err := z.osOpenFn(filePath)
	if err != nil {
		return err
	}
	_, err = z.ioCopyFn(fileInZipWriter, fsFile)
	if err != nil {
		return err
	}
	return nil
}

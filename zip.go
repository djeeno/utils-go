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

func (z *zipT) ArchivesRecursive(zipFilePath string, targetPaths []string, withoutRootDirectory bool) (errForDeferClose error) {
	if zipFilePath == "" || zipFilePath == "/" {
		zipFilePath = "./default.zip"
	}

	zipFile, err := z.osCreateFn(zipFilePath)
	if err != nil {
		return err
	}

	zw := zip.NewWriter(zipFile)
	defer func() {
		if errForDeferClose != nil {
			_ = zw.Close()
			return
		}
		errForDeferClose = zw.Close()
	}()

	for _, targetPath := range targetPaths {
		// resolve path
		path, err := filepath.Abs(targetPath)
		if err != nil {
			return err
		}
		// walkFunc
		walkFunc := func(filePath string, info os.FileInfo, err error) error {
			prefixTrimmingRelativePath := filepath.Dir(path)
			return funcForWalkFunc(z, zw, prefixTrimmingRelativePath, filePath, info, err)
		}
		if withoutRootDirectory {
			walkFunc = func(filePath string, info os.FileInfo, err error) error {
				prefixTrimmingRelativePath := path
				return funcForWalkFunc(z, zw, prefixTrimmingRelativePath, filePath, info, err)
			}
		}
		// walk
		if err := z.filepathWalkFn(path, walkFunc); err != nil {
			return err
		}
	}

	return nil
}

func funcForWalkFunc(z *zipT, zw *zip.Writer, prefixTrimmingRelativePath string, filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if prefixTrimmingRelativePath == filePath {
		if info.IsDir() {
			return nil
		}
		// こうしなければ TrimPrefix() において s == prefix になってしまい、 relativePath == "" になってしまう。
		prefixTrimmingRelativePath = filepath.Dir(filePath)
	}

	// /trim/path/to => /path/to
	trimmedPath := strings.TrimPrefix(filePath, prefixTrimmingRelativePath)

	// /path/to => path/to
	relativePath := strings.TrimPrefix(trimmedPath, "/")

	if info.IsDir() && !strings.HasSuffix(relativePath, "/") {
		// path/to/dir => path/to/dir/
		relativePath = relativePath + "/" // for adding empty directory to archive
	}

	destinationFile, err := z.zipWriterCreateFn(zw, relativePath)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil // for adding empty directory to archive
	}

	sourceFile, err := z.osOpenFn(filePath)
	if err != nil {
		return err
	}

	_, err = z.ioCopyFn(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

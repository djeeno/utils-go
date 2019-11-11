package utils

import (
	"archive/zip"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const (
	testTmpZipDirPath = testTmpRootDirPath + "/zip"
)

func initTest() {
	// script->
	initialScript := `
rm -rf ` + testTmpZipDirPath + `
mkdir -p ` + testTmpZipDirPath + `/dir
touch ` + testTmpZipDirPath + `/file
`
	// <-script
	// Run
	out, err := exec.Command("/bin/bash", "-c", initialScript).CombinedOutput()
	if err != nil {
		Log.Fatalfln("TestZipT_ArchivesRecursive(): %v, %v", strings.TrimSuffix(string(out), "\n"), err)
	}
}

func TestZipT_ArchivesRecursive(t *testing.T) {
	t.Helper()
	initTest()

	t.Run("non-normal/ioCopyFn", func(t *testing.T) {
		failureZIP := ZIP
		failureZIP.ioCopyFn = func(dst io.Writer, src io.Reader) (written int64, err error) {
			return 0, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/test.zip", testTmpZipDirPath, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/osOpenFn", func(t *testing.T) {
		failureZIP := ZIP
		failureZIP.osOpenFn = func(name string) (*os.File, error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/test.zip", testTmpZipDirPath, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/zipWriterCreateFn", func(t *testing.T) {
		failureZIP := ZIP
		failureZIP.zipWriterCreateFn = func(w *zip.Writer, name string) (io.Writer, error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/test.zip", testTmpZipDirPath, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/osCreateFn", func(t *testing.T) {
		failureZIP := ZIP
		failureZIP.osCreateFn = func(name string) (*os.File, error) {
			return nil, ErrorDummyErrorForTest
		}
		if err := failureZIP.ArchivesRecursive(testTmpZipDirPath+"/test.zip", testTmpZipDirPath, false); err == nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err == nil: %v", err)
		}
	})

	t.Run("non-normal/walkFuncForArchivesRecursive", func(t *testing.T) {
		if err := walkFuncForArchivesRecursive(nil, nil, "", "", nil, ErrorDummyErrorForTest); err == nil {
			t.Error("TestZipT_ArchivesRecursive(): walkFuncForArchivesRecursive(): err == nil")
		}
	})

	t.Run("normal/withoutRootDirectory=false", func(t *testing.T) {
		if err := ZIP.ArchivesRecursive(testTmpZipDirPath+"/test.zip", testTmpZipDirPath, false); err != nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err != nil: %v", err)
		}
	})

	t.Run("normal/withoutRootDirectory=true", func(t *testing.T) {
		if err := ZIP.ArchivesRecursive(testTmpZipDirPath+"/test.zip", testTmpZipDirPath, true); err != nil {
			t.Errorf("TestZipT_ArchivesRecursive(): ZIP.ArchivesRecursive(): err != nil: %v", err)
		}
	})
}

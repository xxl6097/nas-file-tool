package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	if err := EnsureDir(filepath.Dir(dst)); err != nil {
		return err
	}

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func CopyFileToDir(src string, dstDir string) error {
	if err := EnsureDir(dstDir); err != nil {
		return err
	}
	// 获取源文件的文件名
	fileName := filepath.Base(src)
	// 拼接目标文件的完整路径
	dst := filepath.Join(dstDir, fileName)

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func EnsureDir(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		// 其他错误
		return err
	}
	// 不存在，创建
	return os.MkdirAll(path, 0755)
}

func MoveFileToDir(src string, dstDir string) error {
	if err := EnsureDir(dstDir); err != nil {
		return err
	}
	// 获取源文件的文件名
	fileName := filepath.Base(src)
	// 拼接目标文件的完整路径
	dst := filepath.Join(dstDir, fileName)
	return os.Rename(src, dst)
}

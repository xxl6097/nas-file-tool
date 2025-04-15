package utils

import (
	"fmt"
	"io"
	"nas-file-tool/pkg/input"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func scanDir(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			fmt.Println("子目录:", path)
			scanDir(path) // 递归处理子目录
		} else {
			fmt.Println("文件:", path)
		}
	}
	return nil
}

// WildcardReplace 函数将包含通配符的模式字符串转换为正则表达式模式，并替换通配符匹配部分
func WildcardReplace(pattern, replacement, input string) string {
	// 将通配符 * 转换为正则表达式的 (.*)，使用捕获组来匹配通配符部分
	regexPattern := strings.ReplaceAll(pattern, "*", "(.*)")
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		fmt.Println("正则表达式编译错误:", err)
		return input
	}
	// 执行替换操作，将捕获组替换为目标字符串
	return re.ReplaceAllString(input, strings.ReplaceAll(pattern, "*", replacement))
}

func RenameFiles(dir, pattern, target string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			fileName := strings.ReplaceAll(entry.Name(), pattern, target)
			fileName = WildcardReplace(pattern, target, fileName)
			fmt.Println("-", fileName)
		}
	}
	if !input.Confirm() {
		return nil
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			fileName := strings.ReplaceAll(entry.Name(), pattern, target)
			fileName = WildcardReplace(pattern, target, fileName)
			//fmt.Println(fileName)
			err = os.Rename(filepath.Join(dir, entry.Name()), filepath.Join(dir, fileName))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

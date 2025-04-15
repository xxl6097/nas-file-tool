package internal

import (
	"fmt"
	"nas-file-tool/internal/menu"
	"nas-file-tool/pkg/input"
	"nas-file-tool/pkg/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Registe() {
	menu.GetInstance().AddMenuItem("移动", filesMove)
	menu.GetInstance().AddMenuItem("复制", filesCopy)
	menu.GetInstance().AddMenuItem("重命名", filesRename)
	menu.GetInstance().AddMenuItem("复制【含子目录】", copyFilesWithAllChildren)
}

func filesMove() {
	pattern := input.InputString("请输入匹配文件名(通配符):")
	srcDir := input.InputString("源目录路径:")
	dstInput := input.InputString("目标路径:")
	matches, _ := filepath.Glob(filepath.Join(srcDir, pattern))

	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(srcDir, fileName)
		fmt.Println("-", srcPath)
	}

	if !input.Confirm("确定移动吗") {
		return
	}

	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(srcDir, fileName)
		go func(srcPath, dstInput string) {
			err := utils.MoveFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("移动失败", srcPath, err)
			}
		}(srcPath, dstInput)
	}
}

func filesCopy() {
	pattern := input.InputString("请输入匹配文件名(通配符):")
	srcDir := input.InputString("源目录路径:")
	dstInput := input.InputString("目标路径：")
	matches, _ := filepath.Glob(filepath.Join(srcDir, pattern))
	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(srcDir, fileName)
		fmt.Println("-", srcPath)
	}

	if !input.Confirm("确定复制吗") {
		return
	}
	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(srcDir, fileName)
		go func(srcPath, dstInput string) {
			err := utils.CopyFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("复制失败", srcPath, err)
			}
		}(srcPath, dstInput)
	}
}

func copyFilesWithAllChildren() {
	//root := "/Users/uuxia/Desktop/work/code/github/golang/go-frp-panel/*.go"
	// /Users/uuxia/Desktop/work/code/local/go/nas-file-tool/cmd/test001
	// /Users/uuxia/Desktop/work/code/github/golang/go-frp-panel/*.go
	//pattern := "*.go" // 支持通配符如 *.txt 或 logs/​**​/*.log

	pattern := input.InputString("请输入匹配文件名(通配符):")
	srcDir := input.InputString("源目录路径:")

	dstInput := input.InputString("目标路径：")

	// 将通配符转换为正则表达式（支持 * 和 ​**​）
	regexPattern := strings.ReplaceAll(pattern, ".", `\.`)
	regexPattern = strings.ReplaceAll(regexPattern, "*", ".*")
	re := regexp.MustCompile("^" + regexPattern + "$")

	var matches []string
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		// 提取相对路径用于匹配（如 data/logs/error.log）
		relPath, _ := filepath.Rel(srcDir, path)
		if re.MatchString(relPath) {
			matches = append(matches, path)
			fmt.Println("-", relPath)
		}
		return nil
	})
	fmt.Println("匹配文件数量:", len(matches))
	if !input.Confirm("确定复制吗") {
		return
	}
	for _, f := range matches {
		//fmt.Println("-", f)
		go func(srcPath, dstInput string) {
			err := utils.CopyFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("复制失败", srcPath, err)
			}
		}(f, dstInput)
	}
}

// /Users/uuxia/Desktop/work/code/github/golang/nas-file-tool/test
func filesRename() {
	srcDir := input.InputString("源路径：")
	words := input.InputString("通配符匹配：")
	keywords := input.InputStringEmpty("替换字符：", "")
	_ = utils.RenameFiles(srcDir, words, keywords)
}

package internal

import (
	"fmt"
	"nas-file-tool/pkg/input"
	"nas-file-tool/pkg/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func filesMove() {
	srcInput := input.InputString("源路径(文件名可用通配符)：")
	dstInput := input.InputString("目标路径：")
	director := filepath.Dir(srcInput)
	matches, _ := filepath.Glob(srcInput)

	if !input.Confirm() {
		return
	}
	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(director, fileName)
		go func(srcPath, dstInput string) {
			err := utils.MoveFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("移动失败", srcPath, err)
			}
		}(srcPath, dstInput)
	}
}

func filesCopy() {
	srcInput := input.InputString("源路径(文件名可用通配符)：")
	dstInput := input.InputString("目标路径：")
	director := filepath.Dir(srcInput)
	matches, _ := filepath.Glob(srcInput)
	if !input.Confirm() {
		return
	}
	for _, path := range matches {
		fileName := filepath.Base(path)
		srcPath := filepath.Join(director, fileName)
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

	srcInput := input.InputString("源路径(文件名可用通配符)：")
	root, pattern := filepath.Split(srcInput)
	dstInput := input.InputString("目标路径：")

	// 将通配符转换为正则表达式（支持 * 和 ​**​）
	regexPattern := strings.ReplaceAll(pattern, ".", `\.`)
	regexPattern = strings.ReplaceAll(regexPattern, "*", ".*")
	re := regexp.MustCompile("^" + regexPattern + "$")

	var matches []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		// 提取相对路径用于匹配（如 data/logs/error.log）
		relPath, _ := filepath.Rel(root, path)
		if re.MatchString(relPath) {
			matches = append(matches, path)
		}
		return nil
	})
	fmt.Println("匹配文件数量:", len(matches))
	if !input.Confirm() {
		return
	}
	for _, f := range matches {
		fmt.Println("-", f)
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

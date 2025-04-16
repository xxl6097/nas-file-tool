package internal

import (
	"fmt"
	"nas-file-tool/internal/menu"
	"nas-file-tool/pkg/input"
	"nas-file-tool/pkg/utils"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	menu.GetInstance().AddMenuItem("移动", filesMove)
	menu.GetInstance().AddMenuItem("复制", filesCopy)
	menu.GetInstance().AddMenuItem("重命名", filesRename)
	menu.GetInstance().AddMenuItem("复制【含子目录】", copyFilesWithAllChildren)
	menu.GetInstance().AddMenuItem("IPTV清单", forIPTV)
}

func Menu() {
	menu.GetInstance().ShowMenu()
}

func forIPTV() {
	rootDir := input.InputString("源目录路径:")
	urls := input.Input("请输入视频源规则:")
	utils.FindMoves(rootDir, strings.Split(urls, " "))
}

func filesMove() {
	pattern := input.InputString("请输入匹配文件名(通配符):")
	srcDir := input.InputString("源目录路径:")
	dstInput := input.InputString("目标路径:")
	utils.Movefiles(pattern, srcDir, dstInput)
}

func filesCopy() {
	pattern := input.InputString("请输入匹配文件名(通配符):")
	srcDir := input.InputString("源目录路径:")
	dstInput := input.InputString("目标路径：")
	utils.CopyFiles(pattern, srcDir, dstInput)

}

func copyFilesWithAllChildren() {
	//root := "/Users/uuxia/Desktop/work/code/github/golang/go-frp-panel/*.go"
	// /Users/uuxia/Desktop/work/code/local/go/nas-file-tool/cmd/test001
	// /Users/uuxia/Desktop/work/code/github/golang/go-frp-panel/*.go
	//pattern := "*.go" // 支持通配符如 *.txt 或 logs/​**​/*.log

	pattern := input.InputString("请输入匹配文件名(通配符):")
	srcDir := input.InputString("源目录路径:")
	dstInput := input.InputString("目标路径：")

	utils.CopyChildrenFiles(pattern, srcDir, dstInput)

}

// /Users/uuxia/Desktop/work/code/github/golang/nas-file-tool/video/电影/其他电影
func filesRename() {
	srcDir := input.InputString("源路径：")
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		isDir := entry.IsDir()
		if !isDir {
			fileName := strings.Replace(entry.Name(), " ", "", -1)
			err = os.Rename(filepath.Join(srcDir, entry.Name()), filepath.Join(srcDir, fileName))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("-", fileName)
		}
	}

	for {
		words := input.InputString("通配符匹配：")
		keywords := input.InputStringEmpty("替换字符：", "")
		_ = utils.RenameFiles(srcDir, words, keywords)
	}
}

func trimSpace() {
	srcDir := input.InputString("源路径：")
	_ = utils.TrimSpace(srcDir)
}

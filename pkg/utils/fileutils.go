package utils

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"nas-file-tool/pkg/input"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func IsVideoFile(filePath string) bool {
	ext := filepath.Ext(filePath)
	ext = strings.ToLower(ext)
	switch ext {
	case ".mp4", ".avi", ".mov", ".wmv", ".flv", ".mkv", ".rmvb":
		return true
	default:
		return false
	}
}

func IsVideoFile1(filePath string) bool {
	ext := filepath.Ext(filePath)
	ext = strings.ToLower(ext)
	switch ext {
	case ".go":
		return true
	default:
		return false
	}
}

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

// 通配符转正则表达式
func wildcardToRegex(pattern string) string {
	pattern = regexp.QuoteMeta(pattern)               // 转义特殊字符
	pattern = strings.ReplaceAll(pattern, `\*`, `.*`) // 替换 * 为 .*
	pattern = strings.ReplaceAll(pattern, `\?`, `.`)  // 替换 ? 为 .
	return "^" + pattern + "$"                        // 限定整个字符串匹配
}

// WildcardReplace 函数将包含通配符的模式字符串转换为正则表达式模式，并替换通配符匹配部分
func WildcardReplace(pattern, replacement, input string) string {
	// 将通配符 * 转换为正则表达式的 (.*)，使用捕获组来匹配通配符部分
	regexPattern := strings.ReplaceAll(pattern, "*", "(.*)")
	regexPattern = wildcardToRegex(regexPattern)
	re, err := regexp.Compile(regexPattern)
	if err != nil {
		fmt.Println("正则表达式编译错误:", err)
		return input
	}
	// 执行替换操作，将捕获组替换为目标字符串
	return re.ReplaceAllString(input, strings.ReplaceAll(pattern, "*", replacement))
}

func isWildcardMatch(str, pattern string) bool {
	regexPattern := wildcardToRegex(pattern)
	re := regexp.MustCompile(regexPattern)
	return re.MatchString(str)
}

func RenameFiles(dir, pattern, target string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		isMatch := isWildcardMatch(entry.Name(), pattern)
		isMatch = isMatch || strings.Contains(entry.Name(), pattern)
		isDir := entry.IsDir()
		//fmt.Println(isMatch, entry.Name(), pattern)
		if !isDir && isMatch {
			fileName := strings.ReplaceAll(entry.Name(), pattern, target)
			fileName = strings.ReplaceAll(fileName, " ", "")
			fileName = WildcardReplace(pattern, target, fileName)
			fmt.Println("-", fileName)
		}
	}
	if !input.Confirm("确定重命名吗") {
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

func TrimSpace(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		isMatch := strings.Contains(entry.Name(), " ")
		isDir := entry.IsDir()
		if !isDir && isMatch {
			fileName := strings.Replace(entry.Name(), " ", "", -1)
			fmt.Println("-", fileName)
		}
	}
	if !input.Confirm("确定重命名吗") {
		return nil
	}
	for _, entry := range entries {
		isMatch := strings.Contains(entry.Name(), " ")
		isDir := entry.IsDir()
		if !isDir && isMatch {
			fileName := strings.Replace(entry.Name(), " ", "", -1)
			err = os.Rename(filepath.Join(dir, entry.Name()), filepath.Join(dir, fileName))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func CopyChildrenFiles(pattern, srcDir, dstInput string) {
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
			err := CopyFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("复制失败", srcPath, err)
			}
		}(f, dstInput)
	}
}

func CopyFiles(pattern, srcDir, dstInput string) {
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
			err := CopyFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("复制失败", srcPath, err)
			}
		}(srcPath, dstInput)
	}
}

func Movefiles(pattern, srcDir, dstInput string) {
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
			err := MoveFileToDir(srcPath, dstInput)
			if err != nil {
				fmt.Println("移动失败", srcPath, err)
			}
		}(srcPath, dstInput)
	}
}

func FindMoves(rootDir string, urls []string) {
	videos := make(map[string][]string, 0)
	_ = filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			//fmt.Println("-", d.Name())
			return nil
		}
		// 提取相对路径用于匹配（如 data/logs/error.log）
		relPath, _ := filepath.Rel(rootDir, path)
		if IsVideoFile(relPath) {
			parentDir := filepath.Dir(filepath.Clean(relPath))
			dir := filepath.Base(parentDir)
			video := videos[dir]
			if video == nil {
				video = make([]string, 0)
			}
			video = append(video, relPath)
			//fmt.Println(" -", relPath)
			videos[dir] = video
		}

		return nil
	})
	///root/z4/18688947359/data/public/video
	//http://uuxia.cn:8086/chfs/shared/video/%s?v=1$!-zzzzzzz http://192.168.1.2:8086/chfs/shared/video/%s?v=1$!-zzzzzzz
	//http://192.168.1.2:8086/chfs/shared/video/%s?v=1$!-zzzzzzz
	sb := strings.Builder{}
	for key, v := range videos {
		sb.WriteString(fmt.Sprintf("🎬%s,#genre#\n", key))
		//fmt.Println(fmt.Sprintf("🎬%s,#genre#", key))
		for _, file := range v {
			filename := filepath.Base(file)
			filename = strings.TrimSuffix(filename, filepath.Ext(filename))
			for i := 0; i < len(urls); i++ {
				u := fmt.Sprintf(urls[i], file)
				str := fmt.Sprintf("%s,%s", filename, u)
				if i == len(urls)-1 {
					sb.WriteString(str)
				} else {
					sb.WriteString(fmt.Sprintf("%s\n", str))
				}

			}
		}

		sb.WriteString("\n")
	}

	content := []byte(sb.String())
	err := os.WriteFile("./output.txt", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

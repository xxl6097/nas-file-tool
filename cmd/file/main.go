package main

import (
	"nas-file-tool/internal"
	"nas-file-tool/internal/menu"
)

func main() {
	internal.Registe()
	menu.GetInstance().ShowMenu()

	//rootDir := "/Users/uuxia/Desktop/work/code/github/golang/go-frp-panel"
	//videos := make(map[string][]string, 0)
	//_ = filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
	//	if err != nil || d.IsDir() {
	//		//fmt.Println("-", d.Name())
	//		return nil
	//	}
	//	// 提取相对路径用于匹配（如 data/logs/error.log）
	//	relPath, _ := filepath.Rel(rootDir, path)
	//	if utils.IsVideoFile(relPath) {
	//		parentDir := filepath.Dir(filepath.Clean(relPath))
	//		dir := filepath.Base(parentDir)
	//		video := videos[dir]
	//		if video == nil {
	//			video = make([]string, 0)
	//		}
	//		video = append(video, relPath)
	//		//fmt.Println(" -", relPath)
	//		videos[dir] = video
	//	}
	//
	//	return nil
	//})
	//for key, v := range videos {
	//	fmt.Println(fmt.Sprintf("🎬%s,#genre#", key))
	//	for _, file := range v {
	//		//fmt.Sprintf("决战中途岛.2019.BD1080p.国英双语.中英双字.mp4,http://uuxia.cn:8086/chfs/shared/video/电影/战争电影/决战中途岛.2019.BD1080p.国英双语.中英双字.mp4?v=1$!-zzzzzzz")
	//		filename := filepath.Base(file)
	//		filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	//		str := fmt.Sprintf("%s,%s/chfs/shared/video/%s?v=1$!-zzzzzzz", filename, "http://uuxia.cn:8086", filename)
	//		//fmt.Println(file)
	//		fmt.Println(str)
	//	}
	//}
}

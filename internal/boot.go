package internal

import "nas-file-tool/internal/menu"

func Registe() {
	menu.GetInstance().AddMenuItem("移动", filesMove)
	menu.GetInstance().AddMenuItem("复制", filesCopy)
	menu.GetInstance().AddMenuItem("重命名", filesRename)
	menu.GetInstance().AddMenuItem("复制【含子目录】", copyFilesWithAllChildren)
}

package root_dir

import (
	"os"
	"path/filepath"
)

// inferRootDir 找到包含subDir的目录
// 一般subDir都在根目录下，而main文件在cmd或其他文件下，通过该函数能够找到项目的根路径。

func inferRootDir(subDir string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if exists(d + sub) {
			return d
		}
		return infer(filepath.Dir(d))
	}
	return infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

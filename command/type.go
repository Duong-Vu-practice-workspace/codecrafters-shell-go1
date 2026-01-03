package command

import (
	"os"
	"path/filepath"
)

func checkPath(cmd string) (string, bool) {
	dirs := filepath.SplitList(os.Getenv("PATH"))
	for _, dir := range dirs {
		fullPath := filepath.Join(dir, cmd)
		info, err := os.Stat(fullPath)
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		//check permission
		if info.Mode()&0111 != 0 {
			return fullPath, true
		}
	}
	return "", false
}

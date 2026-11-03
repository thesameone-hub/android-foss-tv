package git

import (
	"os"
	"os/exec"
)

func CloneRepo(gitUrl string) (dirPath string, err error) {
	dirPath, err = os.MkdirTemp("", "git-*")
	if err != nil {
		return
	}

	// The importer only scans the current source tree for screenshots.
	// A shallow clone avoids downloading the full commit history.
	cloneCmd := exec.Command("git", "clone", "--depth=1", gitUrl, dirPath)
	err = cloneCmd.Run()
	if err != nil {
		_ = os.RemoveAll(dirPath)
		return
	}

	return
}

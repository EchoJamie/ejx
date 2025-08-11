/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package hexo

import (
	"io"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

const groupId = "hexo"
const groupTitle = groupId + " Mode"

func hexoClean() error {
	cleanCmd := exec.Command("hexo", "clean")
	cleanCmd.Stdout = io.Discard
	cleanCmd.Stderr = io.Discard
	cleanCmd.Dir = getBlogRootPath()
	err := cleanCmd.Run()
	return err
}

func hexoGenerate(draft bool) error {
	cmd := []string{"generate"}
	if draft {
		cmd = append(cmd, "--draft")
	}
	generateCmd := exec.Command("hexo", cmd...)
	generateCmd.Stdout = io.Discard
	generateCmd.Stderr = io.Discard
	generateCmd.Dir = getBlogRootPath()
	return generateCmd.Run()
}

func hexoServer(draft bool) error {
	cmd := []string{"server"}
	if draft {
		cmd = append(cmd, "--draft")
	}
	serverCmd := exec.Command("hexo", cmd...)
	serverCmd.Stdout = os.Stdout
	serverCmd.Stderr = io.Discard
	serverCmd.Dir = getBlogRootPath()
	err := serverCmd.Run()
	return err
}

func hexoNewDraft(title string) error {
	newDraftCmd := exec.Command("hexo", "new", "draft", title)
	newDraftCmd.Stdout = io.Discard
	newDraftCmd.Stderr = io.Discard
	newDraftCmd.Dir = getBlogRootPath()
	err := newDraftCmd.Run()
	return err
}

func hexoPublish(name string) error {
	publishCmd := exec.Command("hexo", "publish", name)
	publishCmd.Stdout = io.Discard
	publishCmd.Stderr = io.Discard
	publishCmd.Dir = getBlogRootPath()
	err := publishCmd.Run()
	return err
}

func open() error {
	openCmd := exec.Command("code", ".")
	openCmd.Stdout = io.Discard
	openCmd.Stderr = io.Discard
	openCmd.Dir = getBlogRootPath()
	err := openCmd.Run()
	return err
}

func getBlogRootPath() string {
	return viper.GetString("hexo.root")
}

func extendsBlogRootPath(childPath string) string {
	return getBlogRootPath() + childPath
}

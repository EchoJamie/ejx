/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package main

import (
	// import to initialize global config.
	_ "github.com/EchoJamie/ejx/init/global"
	"github.com/EchoJamie/ejx/pkg/ejx"
	_ "github.com/EchoJamie/ejx/pkg/hexo"
)

func main() {
	ejx.Run()
}

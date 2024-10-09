/*
Copyright © 2024 EchoJamie HERE <EMAIL ADDRESS>
*/
package interaction

import (
	"fmt"
	"github.com/EchoJamie/ejx/tools/text"
	"os"
	"strings"
)

func SecondCheck(checkText string) {
	SecondCheckWithYesAndNo(checkText, "", "")
}

func SecondCheckWithYes(checkText, yes string) {
	SecondCheckWithYesAndNo(checkText, yes, "")
}

func SecondCheckWithNo(checkText, no string) {
	SecondCheckWithYesAndNo(checkText, "", no)
}

func SecondCheckWithYesAndNo(checkText, yes, no string) {
	var yn string
	fmt.Printf("%s %s: ", checkText, text.WhiteText("(Y/N)"))
	defer inputErr()
	fmt.Scanln(&yn)
	yn = strings.ToUpper(yn)
	if yn == "N" {
		fmt.Print(no)
	} else if yn == "Y" {
		fmt.Print(yes)
	} else {
		fmt.Println("输入有误, 请输入Y或N)")
		os.Exit(0)
	}
}

func inputErr() {
	err := recover()
	if err != nil {
		fmt.Println("输入异常, ", err)
		os.Exit(1)
	}
}

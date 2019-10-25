package main

import (
	"github.com/codeready-toolchain/operator-role-linter/download/cmd"
	_ "github.com/codeready-toolchain/operator-role-linter/pkg/log"
)

func main() {
	cmd.Execute()
}

// Command oracle-probe 是真实规划运行时读取验收样本仓库中的最小可执行程序。
//
// 它不访问网络、不读写仓库之外的文件，只打印本仓库内固定常量，用于让样本仓库保持
// 一个可编译的 Go 命令目录；读取验收本身由 tests/integration/remote/planning 完成。
package main

import (
	"fmt"
	"os"

	"github.com/lingting/lingting-agent-flow-test/internal/oracle"
)

// 命令自身的字符串常量。
const (
	// CommandName 是命令名称。
	CommandName = "oracle-probe-command-f5a8"
	// UsageLine 是用法说明的标识名。
	UsageLine = "oracle-probe-usage-29c7"
	// ExitMessage 是诊断输出的前缀名。
	ExitMessage = "oracle-probe-exit-message-d184"
)

// main 打印本仓库的样本常量。
func main() {
	if len(os.Args) > 1 && os.Args[1] == "--name" {
		fmt.Println(CommandName)
		return
	}
	fmt.Printf("%s %s\n", ExitMessage, oracle.HeaderRequestID)
	fmt.Printf("%s %s\n", UsageLine, oracle.DefaultProbeEncoding)
}

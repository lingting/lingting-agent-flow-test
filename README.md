# lingting-agent-flow-test
用于跑测试用例时进行git操作

*手动推送提交测试*

## 读取验收样本

`internal/oracle` 与 `cmd/oracle-probe` 提供真实规划运行时「模型必须读取源码才能回答」的
字符串常量样本：验收只把文件路径与常量标识符交给模型，常量值只能通过读取源码获得。

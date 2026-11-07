// Package oracle 提供真实规划运行时读取验收使用的源码样本。
//
// 该包的字符串常量只作为「模型必须真实读取文件才能回答」的取值样本：调用方只把文件
// 路径与常量标识符交给模型，常量值只能通过读取本文件获得。
package oracle

// 请求与追踪相关的头名称常量。
const (
	// HeaderRequestID 是探针请求标识头。
	HeaderRequestID = "x-oracle-probe-request-8f4c1d2e"
	// HeaderTraceID 是探针链路标识头。
	HeaderTraceID = "x-oracle-probe-trace-5b7a93c6"
	// HeaderSessionID 是探针会话标识头。
	HeaderSessionID = "x-oracle-probe-session-2d7e60ab"
	// HeaderAttemptID 是探针执行尝试标识头。
	HeaderAttemptID = "x-oracle-probe-attempt-c93f1487"
)

// 快照与工作区相关的名称常量。
const (
	// SnapshotVolumePrefix 是快照卷名称前缀。
	SnapshotVolumePrefix = "agentflow-probe-snapshot-b41d"
	// SnapshotWorkDir 是快照卷在容器内的挂载目录名。
	SnapshotWorkDir = "agentflow-probe-workspace-6ea2"
	// SnapshotSessionDir 是快照内的会话目录名。
	SnapshotSessionDir = "agentflow-probe-session-dir-90cf"
	// SnapshotArchiveDir 是快照内的归档目录名。
	SnapshotArchiveDir = "agentflow-probe-archive-dir-7b35"
)

// 规划角色相关的名称常量。
const (
	// RoleRequirementAnalystLabel 是需求分析师角色标签。
	RoleRequirementAnalystLabel = "oracle-role-requirement-analyst-4ae1"
	// RolePlannerLabel 是方案规划师角色标签。
	RolePlannerLabel = "oracle-role-planner-9c6d"
	// RoleInspectorLabel 是巡检角色标签。
	RoleInspectorLabel = "oracle-role-inspector-13f8"
)

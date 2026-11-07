package oracle

// 计划步骤与结果记录的名称常量。
const (
	// StepTypeRequirementAnalysis 是需求分析步骤类型名。
	StepTypeRequirementAnalysis = "oracle-step-requirement-analysis-71bc"
	// StepTypePlanDraft 是方案草稿步骤类型名。
	StepTypePlanDraft = "oracle-step-plan-draft-3e08"
	// StepTypePlanReview 是方案评审步骤类型名。
	StepTypePlanReview = "oracle-step-plan-review-a5d4"
	// StepStatusAwaitingReview 是等待评审状态名。
	StepStatusAwaitingReview = "oracle-status-awaiting-review-6f9a"
	// StepStatusRecorded 是已记录状态名。
	StepStatusRecorded = "oracle-status-recorded-0c2e"
	// RevisionFormatNDJSON 是方案快照的记录格式名。
	RevisionFormatNDJSON = "oracle-format-ndjson-84d1"
	// RevisionFormatSummary 是方案摘要的记录格式名。
	RevisionFormatSummary = "oracle-format-summary-b270"
)

// 探针诊断输出使用的标记常量。
const (
	// DiagnosticStageSnapshot 是快照阶段诊断标记。
	DiagnosticStageSnapshot = "oracle-stage-snapshot-e63c"
	// DiagnosticStageRead 是读取阶段诊断标记。
	DiagnosticStageRead = "oracle-stage-read-1b8f"
	// DiagnosticStageArchive 是归档阶段诊断标记。
	DiagnosticStageArchive = "oracle-stage-archive-77a0"
	// DiagnosticOutcomeAbsent 是不存在结论的诊断标记。
	DiagnosticOutcomeAbsent = "oracle-outcome-absent-58e2"
)

// 探针使用的默认文本常量。
const (
	// DefaultProbePrompt 是探针默认提示词名称。
	DefaultProbePrompt = "oracle-prompt-default-2a6d"
	// DefaultProbeEncoding 是探针默认编码名。
	DefaultProbeEncoding = "oracle-encoding-utf8-c40b"
)

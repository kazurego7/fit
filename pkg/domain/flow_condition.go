package domain

var (
	flowTaskResultStatus = FlowTaskResultStatus{
		completed: "completed",
		failed:    "failed",
		skip:      "skip",
	}
)

type FlowTaskResultStatus struct {
	completed string
	failed    string
	skip      string
}

func NewFlowTaskResultCompleted() FlowTaskResult {
	return FlowTaskResult{status: flowTaskResultStatus.completed}
}

func NewFlowTaskResultFailed() FlowTaskResult {
	return FlowTaskResult{status: flowTaskResultStatus.failed}
}

func NewFlowTaskResultSkipped() FlowTaskResult {
	return FlowTaskResult{status: flowTaskResultStatus.skip}
}

type FlowTaskResult struct {
	status string
}

func (f *FlowTaskResult) IsCompleted() bool {
	return f.status == flowTaskResultStatus.completed
}

func (f *FlowTaskResult) IsFailed() bool {
	return f.status == flowTaskResultStatus.failed
}

func (f *FlowTaskResult) IsSkip() bool {
	return f.status == flowTaskResultStatus.skip
}

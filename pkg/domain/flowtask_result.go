package domain

var (
	FlowTaskResultStatus = flowTaskResultStatus{
		completed: "completed",
		failed:    "failed",
		skip:      "skip",
	}
)

type flowTaskResultStatus struct {
	completed string
	failed    string
	skip      string
}

func (f flowTaskResultStatus) Completed() FlowTaskResult {
	return &flowTaskResult{status: f.completed}
}

func (f flowTaskResultStatus) Failed() FlowTaskResult {
	return &flowTaskResult{status: f.failed}
}

func (f flowTaskResultStatus) Skipped() FlowTaskResult {
	return &flowTaskResult{status: f.skip}
}

type FlowTaskResult interface {
	IsCompleted() bool
	IsFailed() bool
	IsSkip() bool
}

type flowTaskResult struct {
	status string
}

func (f *flowTaskResult) IsCompleted() bool {
	return f.status == FlowTaskResultStatus.completed
}

func (f *flowTaskResult) IsFailed() bool {
	return f.status == FlowTaskResultStatus.failed
}

func (f *flowTaskResult) IsSkip() bool {
	return f.status == FlowTaskResultStatus.skip
}

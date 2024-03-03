package domain

type FlowTask interface {
	Key() string
	Description() string
	CanSkip() bool
	Run() FlowTaskResult
}

type flowTask struct {
	key             string
	description     string
	skipPredicate   func() bool
	skipPresenter   func() string
	taskRunner      func() FlowTaskResult
	rollbackProcess func()
}

func NewFlowTask(key string, description string, skipPredicate func() bool, skipPresenter func() string, taskRunner func() FlowTaskResult, rollbackProcess func()) FlowTask {
	return &flowTask{
		key:             key,
		description:     description,
		skipPredicate:   skipPredicate,
		skipPresenter:   skipPresenter,
		taskRunner:      taskRunner,
		rollbackProcess: rollbackProcess,
	}
}

func (t *flowTask) Key() string {
	return t.key
}

func (t *flowTask) Description() string {
	return t.description
}

func (t *flowTask) CanSkip() bool {
	return t.skipPredicate()
}

func (t *flowTask) Run() FlowTaskResult {
	if t.CanSkip() && t.skipPresenter != nil {
		println(t.skipPresenter())
		return NewFlowTaskResultSkipped()
	}
	result := t.taskRunner()
	if result.IsFailed() && t.rollbackProcess != nil {
		t.rollbackProcess()
	}
	return result
}

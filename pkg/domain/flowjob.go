package domain

import "errors"

type FlowJob interface {
	SetNo(int)
	Key() string
	No() int
	Description() string
	CanShow() bool
	TaskSequence() []FlowTask
	NextRunTaskNo() (int, error)
}

type flowJob struct {
	key              string
	no               int
	description      string
	canShowPredicate func() bool
	taskSequence     []FlowTask
}

func NewFlowJob(key string, description string, canShowPredicate func() bool, taskSequence []FlowTask) FlowJob {
	return &flowJob{
		key:              key,
		description:      description,
		canShowPredicate: canShowPredicate,
		taskSequence:     taskSequence,
	}
}

func (item *flowJob) SetNo(no int) {
	item.no = no
}

func (item flowJob) Key() string {
	return item.key
}

func (item flowJob) No() int {
	return item.no
}

func (item flowJob) Description() string {
	return item.description
}

func (item flowJob) CanShow() bool {
	if item.canShowPredicate == nil {
		return true
	}
	return item.canShowPredicate()
}

func (item flowJob) TaskSequence() []FlowTask {
	return item.taskSequence
}

func (item flowJob) NextRunTaskNo() (int, error) {
	for i, task := range item.taskSequence {
		if !task.CanSkip() {
			return i + 1, nil
		}
	}
	return -1, errors.New("no executable task")
}

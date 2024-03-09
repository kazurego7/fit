package domain

import (
	"errors"
	"sort"

	"github.com/samber/lo"
)

type Workflow interface {
	FilterExecutable(condition FlowJobCondition) Workflow
	ToFlowJobList() []FlowJob
	ExistsFlowJobByNo(no int) bool
	FlowJobByNo(no int) (FlowJob, error)
}

type workflow struct {
	flowJobMap map[string]FlowJob
}

func NewWorkflow(flowJobList []FlowJob) Workflow {
	return workflow{flowJobListToMap(flowJobList)}
}

func (w workflow) FilterExecutable(condition FlowJobCondition) Workflow {
	flowJobList := w.ToFlowJobList()
	lo.ForEach(flowJobList, func(flowJob FlowJob, _ int) {
		flowJob.SetFlowJobCondition(condition)
	})
	filiterdFlowJobList := lo.
		Filter(flowJobList, func(flowJob FlowJob, _ int) bool {
			return flowJob.CanShow()
		})
	return workflow{flowJobListToMap(filiterdFlowJobList)}
}

func flowJobListToMap(flowJobList []FlowJob) map[string]FlowJob {
	var flowJobMap = make(map[string]FlowJob)
	lo.ForEach(flowJobList, func(item FlowJob, index int) {
		item.SetNo(index + 1)
		flowJobMap[item.Key()] = item
	})
	return flowJobMap
}

func (w workflow) ToFlowJobList() []FlowJob {
	var flowJobs = lo.MapToSlice(w.flowJobMap, func(_ string, item FlowJob) FlowJob {
		return item
	})
	sort.SliceStable(flowJobs, func(i, j int) bool {
		return flowJobs[i].No() < flowJobs[j].No()
	})
	return flowJobs
}

func (w workflow) ExistsFlowJobByNo(no int) bool {
	return no >= 1 && no < len(w.flowJobMap)+1
}

func (w workflow) FlowJobByNo(no int) (FlowJob, error) {
	if !w.ExistsFlowJobByNo(no) {
		return nil, errors.New("flowJob not found")
	}
	return w.ToFlowJobList()[no-1], nil
}

package rolexserver

var (
	// collections
	tasks map[string]Task
	agents map[string]Agent
	taskNeeds map[string]TaskNeed
	fulfillments map[string]TaskFulfillment
)

func InitData() {

}

func getTaskByAgentID(agentID string) (Task, error) {
	return Task{}, nil
}

func insertAgent(a Agent)(error) {
	return nil
}

func getAgentByID(agentID string) (Agent, error) {
	return Agent{}, nil
}

func insertTaskNeed(tn TaskNeed) (TaskNeed, error) {
	return TaskNeed{}, nil
}

func insertTask(t Task) (Task, error) {
	return Task{}, nil
}

func getTaskNeedByID(taskNeedID string) (TaskNeed, error) {
	return TaskNeed{}, nil
}

func insertTaskFulfillment(ffm TaskFulfillment)(TaskFulfillment, error) {
	return TaskFulfillment{}, nil
}
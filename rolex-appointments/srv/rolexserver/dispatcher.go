package rolexserver

import (
	"container/list"
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log/slog"
	"sync"
)

func newID(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	// 使用URLEncoding避免+、/等特殊字符
	return base64.URLEncoding.EncodeToString(b)[:length], nil
}


type Dispatcher struct {
	buyers *list.List
	lock sync.Mutex
	store Store
}

func InitDispatcher(buyers []Buyer) {
	// init buyers queue
	dispatcher = Dispatcher{}
	dispatcher.buyers =  list.New()
	for _, b := range buyers {
		dispatcher.buyers.PushBack(b)
	}
}

func (dp *Dispatcher) AssignTask(ctx context.Context,appointmentDate, agentID string) (Task, error) {
	if task, err := dp.store.GetTaskByAgentID(ctx, agentID); err != nil {
		slog.Warn("Already assigned task", "taskid", task.ID, "agentid", agentID)
		return  task, err
	}
	var b Buyer
	var ok bool

	dp.lock.Lock()
	defer dp.lock.Unlock()
	if dp.buyers.Len() == 0 {
		slog.Error("No available buyerbuyer")
		return Task{}, errors.New("no avaialbe buyer")
	}
	front := dp.buyers.Front()
	if b, ok = front.Value.(Buyer); !ok {
		slog.Error("invalid buyer queue")
		return Task{}, errors.New("invalid buyer queue")
	}
	dp.buyers.Remove(front)

	if _id, err := newID(8); err != nil {
		slog.Error("Failed to generate new task ID", "error", err)
		return Task{}, err
	} else {
		task := Task {
			ID:_id,
			AgentID: agentID,
			Param: &TaskParam {
				TargetDate: appointmentDate,
				Buyer: &b,
			},
			State: TaskStateCreated,
		}
		return task, nil
	}
}

func (dp *Dispatcher) RegisterAgent(ctx context.Context, a Agent) (Agent, error) {
	return dp.store.UpsertAgent(ctx, a)
}
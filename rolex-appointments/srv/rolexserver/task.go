package rolexserver

import (
	"container/list"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

var (
	ServerConf *Config
	dispatcher Dispatcher
)

type Buyer struct {
	Agency string `json:"agency" yaml:"agency"`
	Name string `json:"name" yaml:"name"`
	Phone string `json:"phone" yaml:"phone"`
	SSN6 string `json:"ssn6" yaml:"ssn6"`
	SSN1 string `json:"ssn1" yaml:"ssn1"`
}

type Config struct {
	Buyers []Buyer `yaml:"buyers"`
	AppointmentDate string `yaml:"appointment_date"`
}

type Agent struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required,ipv4"`
}

type TaskNeed struct {
	ID string `json:"id"`
	ArtifactB64 string `json:"artifact_b64,omitempty"`
	Type string `json:"type"` // captcha, SMS
}

type TaskFulfillment struct {
	NeedID string `json:"need_id"`
	Type string `json:"type"`
	Text string `json:"text"`
}

type TaskParam struct {
	TargetDate string `json:"target_date,omitempty"`
	Buyer *Buyer `json:"buyer,omitempty"`
}

const (
	TaskStateCreated = "created"
	TaskStateRunning = "running"
	TaskStatePending = "pending"
	TaskStateError = "error"
	TaskStateRestarting = "restarting"
	TaskStateDone = "done"
)

type TaskProgress struct {
	Message string `json:"msg,omitempty"`
	Percent int `json:"percent,omitempty"`
}


type Task struct {
	ID string `json:"id"`
	AgentID string `json:"agent_id"`
	Param *TaskParam `json:"param,omitempty"`
	Progress *TaskProgress`json:"progress,omitempty"`
	State string `json:"state"` // created, running, pending, error, restarting, done
}

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
}

func InitDispatcher(buyers []Buyer) {
	// init buyers queue
	dispatcher = Dispatcher{}
	dispatcher.buyers =  list.New()
	for _, b := range buyers {
		dispatcher.buyers.PushBack(b)
	}
}

func (dp *Dispatcher) AssignTask(appointmentDate, agentID string) (Task, error) {
	if task, err := getTaskByAgentID(agentID); err != nil {
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
		return insertTask(task)
	}
}

func (dp *Dispatcher) RegisterAgent(a Agent) error {
	return insertAgent(a)
}

func LoadConfig(path string)(*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		slog.Error("Failed to load config file", "path", path, "error", err)
		return nil, err
	}
	c, err := io.ReadAll(f)
	if err != nil {
		slog.Error("Failed to read file", "error", err)
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(c, &cfg); err != nil {
		slog.Error("Failed to unmarshal config", "error", err)
		return nil, err
	}
	return &cfg,nil
}

func InitServer(configFilePath string) error {
	cp := ""
	var err error
	if configFilePath == "" {
		cp = "srv.yaml"
	} else {
		cp = configFilePath
	}
	
	if ServerConf, err = LoadConfig(cp); err != nil {
		slog.Error("Failed to init server due to config file error", "error", err)
		return err
	}
	
	InitDispatcher(ServerConf.Buyers)
	InitData()
	return  nil
}

// POST /api/v1/agent
// {
//  "name": "agent-1",
// 	"address":"192.168.1.11"
// }
func RegisterAgentHandler(w http.ResponseWriter,r *http.Request) {
	slog.Info("Registered agent")
	var a Agent
	
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		slog.Error("Failed to read or parse body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var validate = validator.New()
	if err := validate.Struct(&a); err != nil {
		slog.Error("Failed to validate agent info", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := dispatcher.RegisterAgent(a); if err != nil {
		slog.Error("Failed to register agent info", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	if respJson, error :=json.Marshal(a); error != nil {
		w.Write(respJson)
		w.WriteHeader(http.StatusAccepted)
	} else {
		slog.Error("Failed to write response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// POST /api/v1/task?agent=<agent_id>
// Response: task
func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	agentID := r.URL.Query().Get("agent_id")
	slog.Info("Create task", "agentid", agentID)
	if _, err := getAgentByID(agentID); err != nil {
		slog.Error("Unregistered agent", "agent_id", agentID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var task Task
	var err error
	if task, err = dispatcher.AssignTask(ServerConf.AppointmentDate, agentID); err != nil {
		slog.Error("Failed to assign task", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if resp, err := json.Marshal(task); err != nil {
		slog.Error("Failed to marshal task", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(resp)
		w.WriteHeader(http.StatusOK)
	}
}

// POST /api/v1/task/<task-id>/needs
// response needs-id
func PostTaskNeedHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var taskID string
	var ok bool
	if taskID, ok = vars["task-id"]; !ok {
		slog.Error("Failed create task need", "task-id", taskID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	slog.Info("Create task need", "task-id", taskID)
	var need TaskNeed
	if err := json.NewDecoder(r.Body).Decode(&need); err != nil {
		slog.Error("Failed to decode task param", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if tn, err := insertTaskNeed(need); err != nil {
		slog.Error("Failed to insert task need", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	} else {
		if data, err := json.Marshal(tn); err != nil {
			slog.Error("Failed to marshal task need", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.Write(data)
			if _, err := insertTaskFulfillment(TaskFulfillment{
				NeedID: tn.ID,
				Type: tn.Type,
			});err != nil {
				slog.Error("Failed to insert task fulfillment", "error", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
}

// PATCH /api/v1/task/<task-id>/fulfillment/<nee-id>
// user update fulfillment
func PatchTaskFulfillmentHandler(w http.ResponseWriter, r *http.Request) {
}

// GET /api/v1/task/<task-id>/fulfillment/<need-id>
// response fulfillment
func GetTaskFulfillmentHandler(w http.ResponseWriter, r *http.Request) {

}

// Patch /api/v1/task/<id>
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	slog.Info("Put task...")
}
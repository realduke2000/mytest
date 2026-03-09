package rolexserver

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

var (
	ServerConf *Config
	dispatcher Dispatcher
)

const (
	TaskStateCreated = "created"
	TaskStateRunning = "running"
	TaskStatePending = "pending"
	TaskStateError = "error"
	TaskStateRestarting = "restarting"
	TaskStateDone = "done"
)

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
	/*
		Load config
		Clear expired tasks in DB

	*/
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

// PUT /api/v1/agent
// {
//  "name": "agent-1",
// 	"address":"192.168.1.11"
// }
func RegisterAgentHandler(w http.ResponseWriter,r *http.Request) {
	slog.Info("Register agent")
	var a Agent
	if _id, ok := r.Context().Value("agent-id").(string); ok {
		a.ID = _id
	} else {
		slog.Error("Invalid agent-id")
		w.WriteHeader(http.StatusBadRequest)
		return	
	}
	
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
	_, err := dispatcher.RegisterAgent(r.Context(), a); if err != nil {
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
	if _, err := dispatcher.store.GetAgentByID(r.Context(),agentID); err != nil {
		slog.Error("Unregistered agent", "agent_id", agentID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var task Task
	var err error
	if task, err = dispatcher.AssignTask(r.Context(), ServerConf.AppointmentDate, agentID); err != nil {
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
	slog.Info("Creating task need", "task-id", taskID)
	var need TaskNeed
	if err := json.NewDecoder(r.Body).Decode(&need); err != nil {
		slog.Error("Failed to decode task param", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if tn, err := dispatcher.store.InsertTaskNeed(r.Context(), need); err != nil {
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
			if _, err := dispatcher.store.InsertTaskFulfillment(r.Context(), TaskFulfillment{
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

// PUT /api/v1/task/<task-id>/fulfillment/<nee-id>
// user update fulfillment
func PutTaskFulfillmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var taskID string
	var ok bool
	if taskID, ok = vars["task-id"]; !ok {
		slog.Error("Failed create task need", "task-id", taskID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if tn, err := dispatcher.store.TaskFulfillmentsColl().(r.Context(), need); err != nil {
		slog.Error("Failed to update task need", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return 
	} else {
		if data, err := json.Marshal(tn); err != nil {
			slog.Error("Failed to marshal task need", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.Write(data)
			if _, err := dispatcher.store.InsertTaskFulfillment(r.Context(), TaskFulfillment{
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

// GET /api/v1/task/<task-id>/fulfillment/<need-id>
// response fulfillment
func GetTaskFulfillmentHandler(w http.ResponseWriter, r *http.Request) {

}

// Patch /api/v1/task/<id>
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	slog.Info("Put task...")
}
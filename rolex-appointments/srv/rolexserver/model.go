package rolexserver

import "time"

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
    ID      string `json:"id"      bson:"id"`                 // 业务 ID，collection 上唯一索引
    Name    string `json:"name"    bson:"name"    validate:"required"`
    Address string `json:"address" bson:"address" validate:"required,ipv4"`
}

// Dynamic input for a running task
type TaskNeed struct {
    ID          string `json:"id"                     bson:"id"`                       // 业务 ID，唯一
	TaskID      string `json:"task_id" bson:"task_id"`
    ArtifactB64 string `json:"artifact_b64,omitempty" bson:"artifact_b64,omitempty"`   // captcha 图片等
    Type        string `json:"type"                   bson:"type"`                     // captcha, SMS
}

type TaskFulfillment struct {
	ID     string `json:"id"      bson:"id"`       
    TaskID string `json:"task_id" bson:"task_id"`
    NeedID string `json:"need_id" bson:"need_id"`
    Type   string `json:"type"    bson:"type"`
    Text   string `json:"text"    bson:"text"`
}

type TaskParam struct {
    TargetDate string  `json:"target_date,omitempty" bson:"target_date,omitempty"`
    Buyer      *Buyer  `json:"buyer,omitempty"       bson:"buyer,omitempty"`
}

type TaskProgress struct {
    Message string `json:"msg,omitempty"     bson:"msg,omitempty"`
    Percent int    `json:"percent,omitempty" bson:"percent,omitempty"`
}

type Task struct {
    ID       string        `json:"id"              bson:"id"`                  // 业务 ID，唯一
    AgentID  string        `json:"agent_id"        bson:"agent_id"`
    Param    *TaskParam    `json:"param,omitempty" bson:"param,omitempty"`     // static input for task
    Progress *TaskProgress `json:"progress,omitempty" bson:"progress,omitempty"`
    State    string        `json:"state"           bson:"state"`               // created, running, pending, error, restarting, done
	CreatedAt time.Time    `json:"created_at" bson:"created_at"`
	ExpireAt  time.Time    `json:"expire_at"  bson:"expire_at"`
}

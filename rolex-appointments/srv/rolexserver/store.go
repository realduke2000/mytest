package rolexserver

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
// Store 结构体：封装 Mongo 客户端 + DB + 所有数据操作
//

type Store struct {
    client *mongo.Client
    db     *mongo.Database
}

// NewStore 建立 Mongo 连接并返回 Store，同时创建必要的索引
func NewStore(ctx context.Context, uri, dbName string) (*Store, error) {
    cctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(cctx, options.Client().ApplyURI(uri))
    if err != nil {
        slog.Error("mongo connect failed", "err", err, "uri", uri)
        return nil, fmt.Errorf("connect mongo: %w", err)
    }

    if err := client.Ping(cctx, nil); err != nil {
        slog.Error("mongo ping failed", "err", err)
        return nil, fmt.Errorf("ping mongo: %w", err)
    }

    s := &Store{
        client: client,
        db:     client.Database(dbName),
    }

    if err := s.ensureIndexes(context.Background()); err != nil {
        slog.Error("ensure indexes failed", "err", err)
        return nil, err
    }

    slog.Info("mongo connected", "uri", uri, "db", dbName)
    return s, nil
}

func (s *Store) AgentsColl() *mongo.Collection           { return s.db.Collection("agents") }
func (s *Store) TasksColl() *mongo.Collection            { return s.db.Collection("tasks") }
func (s *Store) TaskNeedsColl() *mongo.Collection        { return s.db.Collection("task_needs") }
func (s *Store) TaskFulfillmentsColl() *mongo.Collection { return s.db.Collection("task_fulfillments") }

// ensureIndexes 保证各 collection 的 id 字段有唯一索引，以及常用字段索引
func (s *Store) ensureIndexes(ctx context.Context) error {
    ensureUniqueID := func(coll *mongo.Collection, collName string) error {
        idxModel := mongo.IndexModel{
            Keys: bson.D{{Key: "id", Value: 1}},
            Options: options.Index().SetUnique(true),
        }
        name, err := coll.Indexes().CreateOne(ctx, idxModel)
        if err != nil {
            slog.Error("create unique index on id failed", "collection", collName, "err", err)
            return err
        }
        slog.Info("unique index ensured", "collection", collName, "index", name)
        return nil
    }

    ensureIndex := func(coll *mongo.Collection, collName string, keys bson.D, opts *options.IndexOptions) error {
        idxModel := mongo.IndexModel{
            Keys:    keys,
            Options: opts,
        }
        name, err := coll.Indexes().CreateOne(ctx, idxModel)
        if err != nil {
            slog.Error("create index failed", "collection", collName, "keys", keys, "err", err)
            return err
        }
        slog.Info("index ensured", "collection", collName, "index", name)
        return nil
    }

    // 唯一 id 索引
    if err := ensureUniqueID(s.AgentsColl(), "agents"); err != nil {
        return err
    }
    if err := ensureUniqueID(s.TasksColl(), "tasks"); err != nil {
        return err
    }
    if err := ensureUniqueID(s.TaskNeedsColl(), "task_needs"); err != nil {
        return err
    }

    // 常用查询索引
    if err := ensureIndex(s.TasksColl(), "tasks",
        bson.D{{Key: "agent_id", Value: 1}}, nil); err != nil {
        return err
    }

    if err := ensureIndex(s.TaskNeedsColl(), "task_needs",
        bson.D{{Key: "task_id", Value: 1}}, nil); err != nil {
        return err
    }

    if err := ensureIndex(s.TaskFulfillmentsColl(), "task_fulfillments",
        bson.D{{Key: "task_id", Value: 1}}, nil); err != nil {
        return err
    }

    if err := ensureIndex(s.TaskFulfillmentsColl(), "task_fulfillments",
        bson.D{{Key: "need_id", Value: 1}}, nil); err != nil {
        return err
    }

    // 如果你以后希望任务按过期时间清理，可以加 TTL 索引（这里先不启用）：
    ttlSeconds := int32(7200) // 2 小时
    _ = ensureIndex(s.TasksColl(), "tasks",
        bson.D{{Key: "expire_at", Value: 1}},
        options.Index().SetExpireAfterSeconds(ttlSeconds))

    return nil
}

//
// Agent 相关方法
//

// UpsertAgent 插入一个 Agent，若 ID 为空则自动生成
func (s *Store) UpsertAgent(ctx context.Context, a Agent) (Agent, error) {
    if a.ID == "" {
        a.ID = primitive.NewObjectID().Hex()
    }

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    filter := bson.M{"id": a.ID}
    update := bson.M{"$set":a}
    opts := options.Update().SetUpsert(true)

    _, err := s.AgentsColl().UpdateOne(cctx, filter, update, opts)
    if err != nil {
        slog.Error("UpsertAgent failed", "err", err, "agent", a)
        return Agent{}, err
    }

    slog.Info("UpsertAgent success", "id", a.ID)
    return a, nil
}

// GetAgentByID 根据业务 ID 查询 Agent
func (s *Store) GetAgentByID(ctx context.Context, agentID string) (Agent, error) {
    var a Agent

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    err := s.AgentsColl().FindOne(cctx, bson.M{"id": agentID}).Decode(&a)
    if err != nil {
        slog.Error("GetAgentByID failed", "err", err, "agentID", agentID)
        return Agent{}, err
    }

    slog.Info("GetAgentByID success", "agentID", agentID)
    return a, nil
}

//
// Task 相关方法
//

// InsertTask 插入一个 Task，若 ID / CreatedAt 为空则自动生成
func (s *Store) InsertTask(ctx context.Context, t Task) (Task, error) {
    if t.ID == "" {
        t.ID = primitive.NewObjectID().Hex()
    }
    if t.CreatedAt.IsZero() {
        t.CreatedAt = time.Now().UTC()
    }
    // ExpireAt 由调用方决定；如果你想要默认过期时间，可以在这里设置：

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    _, err := s.TasksColl().InsertOne(cctx, t)
    if err != nil {
        slog.Error("InsertTask failed", "err", err, "task", t)
        return Task{}, err
    }

    slog.Info("InsertTask success", "taskID", t.ID)
    return t, nil
}

// GetTaskByAgentID 查找某个 Agent 的一个 Task（如需未完成任务可加 state 条件）
func (s *Store) GetTaskByAgentID(ctx context.Context, agentID string) (Task, error) {
    var t Task

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    filter := bson.M{
        "agent_id": agentID,
        // 可以按需加状态过滤：
        // "state": bson.M{"$ne": "done"},
    }

    err := s.TasksColl().FindOne(cctx, filter).Decode(&t)
    if err != nil {
        slog.Error("GetTaskByAgentID failed", "err", err, "agentID", agentID)
        return Task{}, err
    }

    slog.Info("GetTaskByAgentID success", "agentID", agentID, "taskID", t.ID)
    return t, nil
}

//
// TaskNeed 相关方法
//

// InsertTaskNeed 插入一个 TaskNeed，若 ID 为空则自动生成，要求 TaskID 必须有
func (s *Store) InsertTaskNeed(ctx context.Context, tn TaskNeed) (TaskNeed, error) {
    if tn.TaskID == "" {
        slog.Error("InsertTaskNeed missing TaskID", "need", tn)
        return TaskNeed{}, fmt.Errorf("task_id is required for TaskNeed")
    }
    if tn.ID == "" {
        tn.ID = primitive.NewObjectID().Hex()
    }

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    _, err := s.TaskNeedsColl().InsertOne(cctx, tn)
    if err != nil {
        slog.Error("InsertTaskNeed failed", "err", err, "need", tn)
        return TaskNeed{}, err
    }

    slog.Info("InsertTaskNeed success", "needID", tn.ID, "taskID", tn.TaskID)
    return tn, nil
}

// GetTaskNeedByID 根据业务 ID 查询 TaskNeed
func (s *Store) GetTaskNeedByID(ctx context.Context, taskNeedID string) (TaskNeed, error) {
    var tn TaskNeed

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    err := s.TaskNeedsColl().FindOne(cctx, bson.M{"id": taskNeedID}).Decode(&tn)
    if err != nil {
        slog.Error("GetTaskNeedByID failed", "err", err, "taskNeedID", taskNeedID)
        return TaskNeed{}, err
    }

    slog.Info("GetTaskNeedByID success", "taskNeedID", taskNeedID, "taskID", tn.TaskID)
    return tn, nil
}

//
// TaskFulfillment 相关方法
//

// InsertTaskFulfillment 插入一个 TaskFulfillment，要求 TaskID 必须有
func (s *Store) InsertTaskFulfillment(ctx context.Context, ffm TaskFulfillment) (TaskFulfillment, error) {
    if ffm.TaskID == "" {
        slog.Error("InsertTaskFulfillment missing TaskID", "fulfillment", ffm)
        return TaskFulfillment{}, fmt.Errorf("task_id is required for TaskFulfillment")
    }

    cctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    _, err := s.TaskFulfillmentsColl().InsertOne(cctx, ffm)
    if err != nil {
        slog.Error("InsertTaskFulfillment failed", "err", err, "fulfillment", ffm)
        return TaskFulfillment{}, err
    }

    slog.Info("InsertTaskFulfillment success", "taskID", ffm.TaskID, "needID", ffm.NeedID)
    return ffm, nil
}

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"
	DB "curiiculum/db"
	"curiiculum/graph/model"
	"curiiculum/models"
	"log"
	"strconv"
	"time"

	null "github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	db := DB.GetDB()
	ex := null.NewString(input.Explanation, input.Explanation != "")
	t, err := time.Parse(time.RFC3339, input.Limit)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// まずTaskテーブルにinsert
	newTask := models.Task{
		ID:          input.ID,
		Title:       input.Title,
		Explanation: ex,
		Limit:       t,
		Priority:    input.Priority,
		Status:      input.Status.String(),
		UserID:      input.UserID,
	}
	if err := newTask.Insert(ctx, db, boil.Infer()); err != nil {
		log.Print(err)
		return nil, err
	}
	// insertしたタスクをselect
	createdTask, err := models.Tasks(qm.Where("id=?", input.ID)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// LabelテーブルからlableIDをselect
	labelID, err := models.Labels(qm.Where("name=?", input.Label)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// task-labelテーブルへのinsert
	newTaskLabel := models.TaskLabelRelation{
		TaskID:  createdTask.ID,
		LabelID: labelID.ID,
	}
	if err := newTaskLabel.Insert(ctx, db, boil.Infer()); err != nil {
		log.Print(err)
		return nil, err
	}
	// task-labelテーブルからlabelIDをselect
	expra := createdTask.Explanation.String

	return &model.Task{
		ID:          createdTask.ID,
		Title:       createdTask.Title,
		Explanation: expra,
		Limit:       createdTask.Limit.Format("2006-01-02"),
		Priority:    createdTask.Priority,
		Status:      model.TaskStatus(createdTask.Status),
		UserID:      createdTask.UserID,
		Label:       input.Label,
	}, nil
}

// UpdateTask is the resolver for the updateTask field.

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	db := DB.GetDB()
	// まずinputを整形
	id := input.ID
	ex := null.NewString(input.Explanation, input.Explanation != "")
	t, err := time.Parse(time.RFC3339, input.Limit)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// updateColumnsを書く
	updateColumns := models.M{
		"ID":          input.ID,
		"Title":       input.Title,
		"Explanation": ex,
		"Limit":       t,
		"Priority":    input.Priority,
		"Status":      input.Status.String(),
		"UserID":      input.UserID,
	}
	// タスクを更新
	_, err = models.Tasks(qm.Where("id = ?", id)).UpdateAll(ctx, db, updateColumns)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	newUpdateTask, err := models.Tasks(qm.Where("id=?", input.ID)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// 新しいラベルのIDを取得
	labelID, err := models.Labels(qm.Where("name=?", input.Label)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// task-labelテーブルも更新
	updateColumns = models.M{
		"task_id":  newUpdateTask.ID,
		"label_id": labelID.ID,
	}
	_, err = models.TaskLabelRelations(qm.Where("task_id=?", newUpdateTask.ID)).UpdateAll(ctx, db, updateColumns)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &model.Task{
		ID:          newUpdateTask.ID,
		Title:       newUpdateTask.Title,
		Explanation: newUpdateTask.Explanation.String,
		Limit:       newUpdateTask.Limit.Format("2006-01-02"),
		Priority:    newUpdateTask.Priority,
		Status:      model.TaskStatus(newUpdateTask.Status),
		UserID:      newUpdateTask.UserID,
		Label:       input.Label,
	}, nil
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, id string) (*model.Task, error) {
	db := DB.GetDB()
	// taskテーブルからselect
	task, err := models.Tasks(qm.Where("id=?", id)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// task-labelテーブルからselect
	tasklabel, err := models.TaskLabelRelations(qm.Where("task_id=?", task.ID)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// labelテーブルからselect
	label, err := models.Labels(qm.Where("id=?", tasklabel.LabelID)).One(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// task-labelテーブルからdelete
	_, err = models.TaskLabelRelations(qm.Where("task_id=?", id)).DeleteAll(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	// taskテーブルからdelete
	_, err = models.Tasks(qm.Where("id=?", id)).DeleteAll(ctx, db)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	lab, err := strconv.Atoi(label.Name)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &model.Task{
		ID:          task.ID,
		Title:       task.Title,
		Explanation: task.Explanation.String,
		Limit:       task.Limit.Format("2006-01-02"),
		Priority:    task.Priority,
		Status:      model.TaskStatus(task.Status),
		UserID:      task.UserID,
		Label:       lab,
	}, nil
}

// User is the resolver for the user field.
// このクエリはやらない
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Print("このクエリは無効です")
	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

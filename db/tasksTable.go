package db

import (
	"context"
	"curiiculum/graph/model"
	"curiiculum/models"
	"time"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func insertIntoTasks(ctx context.Context, input model.NewTask) error {
	db := GetDB()
	t, err := time.Parse(time.RFC3339, input.Limit)
	if err != nil {
		return err
	}
	newTask := models.Task{
		ID:          input.ID,
		Title:       input.Title,
		Explanation: null.NewString(input.Explanation, input.Explanation != ""),
		Limit:       t,
		Priority:    input.Priority,
		Status:      input.Status.String(),
		UserID:      input.UserID,
	}
	if err := newTask.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func updateTasks(ctx context.Context, input model.NewTask) error {
	db := GetDB()
	t, err := time.Parse(time.RFC3339, input.Limit)
	if err != nil {
		return err
	}
	updateColumns := models.M{
		"ID":          input.ID,
		"Title":       input.Title,
		"Explanation": null.NewString(input.Explanation, input.Explanation != ""),
		"Limit":       t,
		"Priority":    input.Priority,
		"Status":      input.Status.String(),
		"UserID":      input.UserID,
	}
	_, err = models.Tasks(qm.Where("id = ?", input.ID)).UpdateAll(ctx, db, updateColumns)
	if err != nil {
		return err
	}
	return nil
}
func deleteFromTask(ctx context.Context, input model.NewTask) error {
	db := GetDB()
	_, err := models.Tasks(qm.Where("id=?", input.ID)).DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

package db

import (
	"context"
	"curiiculum/graph/model"
	"curiiculum/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func insertIntoTaskLabelRelations(ctx context.Context, input model.NewTask, labelID string) error {
	db := GetDB()
	newTaskLabel := models.TaskLabelRelation{
		TaskID:  input.ID,
		LabelID: labelID,
	}
	if err := newTaskLabel.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func updateTaskLabelRelations(ctx context.Context, input model.NewTask, labelID string) error {
	db := GetDB()
	updateColumns := models.M{
		"task_id":  input.ID,
		"label_id": labelID,
	}
	_, err := models.TaskLabelRelations(qm.Where("task_id=?", input.ID)).UpdateAll(ctx, db, updateColumns)
	if err != nil {
		return err
	}
	return nil
}

func deleteTaskLabelRealations(ctx context.Context, input model.NewTask) error {
	db := GetDB()
	_, err := models.TaskLabelRelations(qm.Where("task_id=?", input.ID)).DeleteAll(ctx, db)
	if err != nil {
		return err
	}
	return nil
}

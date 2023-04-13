package db

import (
	"context"
	"curiiculum/graph/model"
)

// 3つのリゾルバはここにつながる。
func CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	err := insertIntoTasks(ctx, input)
	if err != nil {
		return nil, err
	}
	labelID, err := selectFromLabels(ctx, input.Label)
	if err != nil {
		return nil, err
	}
	err = insertIntoTaskLabelRelations(ctx, input, labelID)
	if err != nil {
		return nil, err
	}
	return shapeInput(input), nil
}

func UpdateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	updateTasks(ctx, input)
	labelID, err := selectFromLabels(ctx, input.Label)
	if err != nil {
		return nil, err
	}
	err = updateTaskLabelRelations(ctx, input, labelID)
	if err != nil {
		return nil, err
	}
	return shapeInput(input), nil
}

func DeleteTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	err := deleteTaskLabelRealations(ctx, input)
	if err != nil {
		return nil, err
	}
	err = deleteFromTask(ctx, input)
	if err != nil {
		return nil, err
	}
	return shapeInput(input), nil
}

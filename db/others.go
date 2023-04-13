package db

import (
	"context"
	"curiiculum/graph/model"
	"curiiculum/models"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func shapeInput(input model.NewTask) *model.Task {
	return &model.Task{
		ID:          input.ID,
		Title:       input.Title,
		Explanation: input.Explanation,
		Limit:       input.Limit,
		Priority:    input.Priority,
		Status:      input.Status,
		UserID:      input.UserID,
		Label:       input.Label,
	}
}
func selectFromLabels(ctx context.Context, labelName int) (string, error) {
	db := GetDB()
	labelID, err := models.Labels(qm.Where("name=?", labelName)).One(ctx, db)
	return labelID.ID, err
}

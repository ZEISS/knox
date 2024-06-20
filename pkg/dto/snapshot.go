package dto

import (
	"github.com/zeiss/knox/internal/controllers"
	"github.com/zeiss/knox/internal/models"
	openapi "github.com/zeiss/knox/pkg/apis"
	"github.com/zeiss/knox/pkg/utils"
)

// FromCreateSnapshotRequestObject ...
func FromCreateSnapshotRequestObject(req openapi.CreateSnapshotRequestObject) controllers.CreateSnapshotCommand {
	return controllers.CreateSnapshotCommand{
		Title:       utils.PtrStr(req.Body.Title),
		Description: utils.PtrStr(req.Body.Description),
		StateID:     utils.PtrUUID(req.Body.StateId),
	}
}

// ToCreateSnapshotResponseObject ...
func ToCreateSnapshotResponseObject(snapshot models.Snapshot) openapi.CreateSnapshotResponseObject {
	res := openapi.CreateSnapshot201JSONResponse{}

	return res
}

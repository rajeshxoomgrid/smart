package immhttpservice

import (
	"context"
	"database/sql"

	immdto "github.com/rajeshbond/smart/internal/http/imm/imm_dto"
)

func (s *ImmHttpService) GetMoldChangeStatus(
	ctx context.Context,
	tenantID int64,
	deviceID string,
) (*immdto.MoldChangeStatusResponse, error) {

	command, err := s.ImmHttpService.GetLatestCommand(
		ctx,
		tenantID,
		deviceID,
	)

	if err != nil {
		return nil, err
	}

	if command == nil {
		return nil, sql.ErrNoRows
	}

	response := &immdto.MoldChangeStatusResponse{
		CommandID:       command.ID,
		DeviceID:        command.DeviceID,
		MachineID:       command.MachineID,
		RequestedMoldNo: command.RequestedMoldNo,
		Status:          command.Status,
		CreatedAt:       command.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:       command.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if command.OldMoldNo != "" {
		value := command.OldMoldNo
		response.OldMoldNo = &value
	}

	if command.FailureReason != "" {
		value := command.FailureReason
		response.FailureReason = &value
	}

	if command.ConfirmedAt.IsZero() {
		value := command.ConfirmedAt.Format("2006-01-02T15:04:05Z07:00")
		response.ConfirmedAt = &value
	}

	return response, nil
}

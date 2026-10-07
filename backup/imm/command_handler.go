package immhttp

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	dto "github.com/rajeshbond/smart/internal/mqtt/imm/dto"
	"github.com/rajeshbond/smart/internal/mqtt/imm/service"
)

// ============================================================
// COMMAND HANDLER
// ============================================================

type CommandHandler struct {
	commandService *service.CommandService
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewCommandHandler(
	commandService *service.CommandService,
) *CommandHandler {

	return &CommandHandler{
		commandService: commandService,
	}
}

// ============================================================
// POST COMMAND
// ============================================================
//
// POST
// /api/v1/imm/devices/{deviceID}/command
//
// Example:
//
// /api/v1/imm/devices/04@xoom/command
//
// Body:
//
// {
//     "command": "RESET_COUNTER"
// }
//
// ============================================================

func (h *CommandHandler) SendCommand(
	w http.ResponseWriter,
	r *http.Request,
) {

	// --------------------------------------------------------
	// GET DEVICE ID FROM URL
	// --------------------------------------------------------

	deviceID := chi.URLParam(
		r,
		"deviceID",
	)

	if deviceID == "" {

		writeJSON(
			w,
			http.StatusBadRequest,
			dto.IMMCommandResponse{
				Success: false,
				Message: "device_id is required",
			},
		)

		return
	}

	// --------------------------------------------------------
	// DECODE JSON BODY
	// --------------------------------------------------------

	var req dto.IMMCommandRequest

	if err := json.NewDecoder(
		r.Body,
	).Decode(&req); err != nil {

		writeJSON(
			w,
			http.StatusBadRequest,
			dto.IMMCommandResponse{
				Success:  false,
				Message:  "invalid JSON request",
				DeviceID: deviceID,
			},
		)

		return
	}

	// --------------------------------------------------------
	// PUBLISH MQTT COMMAND
	// --------------------------------------------------------

	topic, err :=
		h.commandService.PublishCommand(
			deviceID,
			req,
		)

	if err != nil {

		log.Printf(
			"❌ IMM COMMAND FAILED | Device=%s | Error=%v",
			deviceID,
			err,
		)

		writeJSON(
			w,
			http.StatusBadGateway,
			dto.IMMCommandResponse{
				Success:  false,
				Message:  err.Error(),
				DeviceID: deviceID,
				Topic:    topic,
			},
		)

		return
	}

	// --------------------------------------------------------
	// SUCCESS RESPONSE
	// --------------------------------------------------------

	writeJSON(
		w,
		http.StatusOK,
		dto.IMMCommandResponse{
			Success:  true,
			Message:  "Command published successfully",
			DeviceID: deviceID,
			Topic:    topic,
		},
	)
}

// ============================================================
// WRITE JSON
// ============================================================

func writeJSON(
	w http.ResponseWriter,
	status int,
	data interface{},
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	if err := json.NewEncoder(
		w,
	).Encode(data); err != nil {

		log.Printf(
			"❌ JSON response error: %v",
			err,
		)
	}
}

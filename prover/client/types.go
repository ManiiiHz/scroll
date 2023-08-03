package client

import (
	"scroll-tech/common/types/message"
)

// RandomResponse defines the response structure for random API
type RandomResponse struct {
	Challenge string `json:"challenge"`
}

// LoginRequest defines the request structure for login API
type LoginRequest struct {
	Message message.AuthMsg `json:"message"`
}

// LoginResponse defines the response structure for login API
type LoginResponse struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
	Data    *struct {
		Time  string `json:"time"`
		Token string `json:"token"`
	} `json:"data,omitempty"`
}

// GetTaskRequest defines the request structure for GetTask API
type GetTaskRequest struct {
	ProverVersion string            `json:"prover_version"`
	ProverHeight  uint64            `json:"prover_height"`
	TaskType      message.ProofType `json:"task_type"`
}

// GetTaskResponse defines the response structure for GetTask API
type GetTaskResponse struct {
	ErrCode int             `json:"errcode,omitempty"`
	ErrMsg  string          `json:"errmsg,omitempty"`
	Data    message.TaskMsg `json:"data,omitempty"`
}

// SubmitProofRequest defines the request structure for the SubmitProof API.
type SubmitProofRequest struct {
	Message message.ProofDetail `json:"message"`
}

// SubmitProofResponse defines the response structure for the SubmitProof API.
type SubmitProofResponse struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

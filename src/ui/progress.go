package ui

import (
	"sync"
)

// TransferState represents the current state of a croc transfer.
// This is used to bridge the blocking croc logic with the Fyne UI.
type TransferState struct {
	mu          sync.RWMutex
	Status      string // "idle" | "connecting" | "transferring" | "done" | "error"
	FilesTotal  int
	FilesDone   int
	BytesTotal  int64
	BytesDone   int64
	RateBps     int64
	CurrentFile string
	Code        string
	Err         error
}

// NewTransferState creates a new TransferState
func NewTransferState() *TransferState {
	return &TransferState{
		Status: "idle",
	}
}

// GetStatus safely returns the current status
func (s *TransferState) GetStatus() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Status
}

// SetStatus safely updates the current status
func (s *TransferState) SetStatus(status string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Status = status
}

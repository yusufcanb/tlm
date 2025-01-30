package shell

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

const checkpointFilename = ".tlm_checkpoint"

// Checkpoint is a struct that represents the last time user checked for updates.
type Checkpoint struct {
	Message        string    `json:"message"`
	Version        string    `json:"version"`
	LastCheckpoint time.Time `json:"time"`
}

// WriteCheckpoint writes the Checkpoint to .tlm_checkpoint in the user's home directory.
func WriteCheckpoint(cp *Checkpoint) error {
	homeDir, err := getHomeDir()
	if err != nil {
		return ShellReadingUserDirErr
	}

	checkpointPath := filepath.Join(homeDir, checkpointFilename)

	data, err := json.Marshal(cp)
	if err != nil {
		return CheckpointFileUnmarshalErr
	}

	// Write the JSON data to the file
	cp.Version = Version
	err = os.WriteFile(checkpointPath, data, 0644)
	if err != nil {
		return CheckpointFileWriteErr
	}

	return nil
}

// GetCheckpoint reads the Checkpoint from .tlm_checkpoint in the user's home directory.
func GetCheckpoint() (*Checkpoint, error) {
	// Get the user's home directory
	homeDir, err := getHomeDir()
	if err != nil {
		return nil, ShellReadingUserDirErr
	}

	// Create the checkpoint file path
	checkpointPath := filepath.Join(homeDir, checkpointFilename)

	// Read the JSON data from the file
	data, err := os.ReadFile(checkpointPath)
	if err != nil {
		return nil, CheckpointFileNotExistErr
	}

	// Unmarshal the JSON data into a time.Time object
	cp := &Checkpoint{}
	err = json.Unmarshal(data, &cp)
	if err != nil {
		return nil, CheckpointFileUnmarshalErr
	}

	return cp, nil
}

func getHomeDir() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return dir, nil
}

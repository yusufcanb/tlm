package shell

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

var CheckpointFileNotExistErr = errors.New("checkpoint file does not exist")
var CheckpointFileUnmarshalErr = errors.New("error unmarshaling checkpoint data")
var CheckpointFileWriteErr = errors.New("error writing checkpoint file")

type Checkpoint struct {
	Message        string    `json:"message"`
	LastCheckpoint time.Time `json:"time"`
}

const checkpointFilename = ".tlm_checkpoint"

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
	err = os.WriteFile(checkpointPath, data, 0644)
	if err != nil {
		return CheckpointFileWriteErr
	}

	return nil
}

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

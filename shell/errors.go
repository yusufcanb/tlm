package shell

import "errors"

var ShellReadingUserDirErr = errors.New("error getting user's home directory")

var CheckpointFileNotExistErr = errors.New("checkpoint file does not exist")
var CheckpointFileUnmarshalErr = errors.New("error unmarshaling checkpoint data")
var CheckpointFileWriteErr = errors.New("error writing checkpoint file")

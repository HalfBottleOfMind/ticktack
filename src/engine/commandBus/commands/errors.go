package commands

import "errors"

var ErrGameNotInProgress = errors.New("game not in progress")
var ErrGameAlreadyInProgress = errors.New("game already in progress")
var ErrInvalidTarget = errors.New("invalid target")

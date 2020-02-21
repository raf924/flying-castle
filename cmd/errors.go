package cmd

import "errors"

var FileNotFoundError = errors.New("the system cannot find the path specified")
var UnreadableFileError = errors.New("cannot read file")
var NotCreatableError = errors.New("cannot create output dir")

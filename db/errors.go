package db

import "errors"

var TransactionError = errors.New("error while committing transaction")
var CreationError = errors.New("error while creating resource")
var ConnectionError = errors.New("error connecting to DB")

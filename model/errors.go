package model

import "errors"

var FileNotFound = errors.New("file doesn't exist")
var WrongFileKind = errors.New("can't use this kind of file here")
var DatabaseError = errors.New("internal server error")
var ContentError = errors.New("could not get file content")
var InvalidNewUsername = errors.New("invalid username value")
var InvalidNewPassword = errors.New("invalid password value")
var SaveFileError = errors.New("cannot save file")
var InvalidCredentials = errors.New("invalid username or password")

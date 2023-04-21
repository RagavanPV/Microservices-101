package model

import (
	"io"
)

type LogOptions struct {
	Level         string
	Writer        io.Writer
}
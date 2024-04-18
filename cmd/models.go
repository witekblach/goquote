package cmd

import (
	"os"
)

type Storage struct {
	Path string
	File *os.File
}

type Quote struct {
	Text   string `json:"text"`
	Viewed int    `json:"viewed"`
}

type Quotes = []Quote

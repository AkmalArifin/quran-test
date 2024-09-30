package models

import "github.com/guregu/null/v5"

type Word struct {
	WordIndex null.Int64  `json:"word_index"`
	WordKey   null.String `json:"word_key"`
	Text      null.String `json:"text"`
}

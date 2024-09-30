package models

import "github.com/guregu/null/v5"

type Page struct {
	PageNumber  null.Int64  `json:"page_number"`
	LineNumber  null.Int64  `json:"line_number"`
	LineType    null.String `json:"line_type"`
	IsCentered  null.Int    `json:"is_centered"`
	FirstWordID null.Int64  `json:"first_word_id"`
	LastWordID  null.Int64  `json:"last_word_id"`
	SurahNumber null.Int64  `json:"surah_number"`
}

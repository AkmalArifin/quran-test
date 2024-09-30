package models

import "github.com/guregu/null/v5"

type Surah struct {
	ID              int64       `json:"id"`
	Name            null.String `json:"name"`
	Arabic          null.String `json:"arabic"`
	TranslationEn   null.String `json:"translation_en"`
	TranslationId   null.String `json:"translation_id"`
	RevelationPlace null.String `json:"revelation_place"`
	RevelationOrder null.String `json:"revelation_order"`
	VersesCount     null.Int    `json:"verses_count"`
	FirstPageID     null.Int    `json:"first_page_id"`
	LastPageID      null.Int    `json:"last_page_id"`
}

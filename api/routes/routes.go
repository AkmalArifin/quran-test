package routes

import (
	"net/http"
	"strconv"

	"example.com/quran/db"
	"example.com/quran/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.GET("/pages/:pageID/:lineID", getPage)
	r.GET("/words/:id", getWord)
	r.GET("/surahs/:id", getSurah)

}

func getPage(c *gin.Context) {
	pageID, err := strconv.ParseInt(c.Param("pageID"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	lineID, err := strconv.ParseInt(c.Param("lineID"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	query := "SELECT * FROM pages WHERE page_number = ? AND line_number = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data.", "error": err.Error()})
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(pageID, lineID)

	var page models.Page
	err = row.Scan(&page.PageNumber, &page.LineNumber, &page.LineType, &page.IsCentered, &page.FirstWordID, &page.LastWordID, &page.SurahNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, page)
}

func getWord(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	query := "SELECT * FROM words WHERE word_index = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data."})
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var word models.Word
	err = row.Scan(&word.WordIndex, &word.WordKey, &word.Text)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data."})
		return
	}

	c.JSON(http.StatusOK, word)
}

func getSurah(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	query := "SELECT * FROM surahs WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data."})
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var surah models.Surah
	err = row.Scan(&surah.ID, &surah.Name, &surah.Arabic, &surah.TranslationEn, &surah.TranslationId, &surah.RevelationPlace, &surah.RevelationOrder, &surah.VersesCount, &surah.FirstPageID, &surah.LastPageID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data."})
		return
	}

	c.JSON(http.StatusOK, surah)
}

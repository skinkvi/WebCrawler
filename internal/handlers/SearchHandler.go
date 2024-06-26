package handlers

import (
	"fmt"
	"net/http"
	"webCrawler/internal/db"
	model "webCrawler/internal/models"

	"github.com/gin-gonic/gin"
)

func SearchHandelerForWebCrawler(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
		return
	}

	var form struct {
		Query string `form:"query"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	query := form.Query

	fmt.Println("Query:", query)

	tx, err := db.DB.Begin()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("Query:", query)

	rows, err := tx.Query("SELECT title, url, text FROM pages WHERE title ILIKE $1 OR text ILIKE $1", "%"+query+"%")

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer rows.Close()

	var results []model.Result
	if rows == nil {
		fmt.Println("No results found")
	} else {
		for rows.Next() {
			var result model.Result
			err := rows.Scan(&result.Title, &result.URL, &result.Text)
			if err != nil {
				fmt.Println("Error scanning row:", err)
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			results = append(results, result)
		}
		fmt.Println("Results:", results)
	}

	fmt.Println("Query:", query)

	err = tx.Commit()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Main website",
		"results": results,
	})

}

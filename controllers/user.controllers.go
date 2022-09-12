package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

var dbConnect *sql.DB

func InitiateDB(db *sql.DB) {
	dbConnect = db
}

type User struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

// get user by id
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := &User{Id: id}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	query := psql.Select("id", "email", "password", "active", "created_at").From("users").Where("id = ?", id)
	sql, args, err := query.ToSql()
	if err != nil {
		log.Println(err)
		return
	}

	rows, err := dbConnect.Query(sql, args...)
	if err != nil {
		log.Printf(" error occured during insertion %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Password, &user.Active, &user.CreatedAt)

		if err != nil {
			log.Printf("error occured during scanning the row %v", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": user,
	})

}

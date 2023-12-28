package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const (
	timeLayout = "2006-01-02 15:04"
)

type Comment struct {
	Id           int       `json:"Id"`
	Descriptions string    `json:"Descriptions"`
	Created_at   time.Time `json:"-"`
	User_id      int       `json:"User_id"`
	TimeFormat   string    `json:"TimeFormat"`
}

func main() {
	connStr := "host=localhost port=5432 user=farrux password=1 dbname=instagram"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	// row := db.QueryRow("SELECT COUNT(1) FROM comment")
	// var count int

	// err = row.Scan(&count)
	// fmt.Println(count, err)
	// var c Comment
	// row = db.QueryRow("select id,descriptions,created_at,user_id from comment")

	// err = row.Scan(
	// 	&c.Id,
	// 	&c.Descriptions,
	// 	&c.Created_at,
	// 	&c.User_id,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// c.TimeFormat = c.Created_at.Format("2006-01-02 15:04:40")
	// fmt.Println(c)
	// c = Comment{
	// 	Descriptions: "Vazifalarni qachonga qilasizlar akalar? Keyingi darsda vazifa qilmaganlar turib eshitadi...",
	// 	User_id:      1,
	// }
	// _, err = db.Exec("INSERT INTO comment (descriptions,user_id) VALUES ($1,$2)",
	// 	c.Descriptions, c.User_id,
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	var comments []Comment

	// data, _ := os.ReadFile("a.json")
	// err = json.Unmarshal(data, &comments)
	// if err != nil {
	// 	fmt.Println("bu xatolik json unmarshal", err)
	// 	return
	// }
	// for _, c := range comments {
	// 	_, _ = db.Exec("INSERT INTO comment (descriptions,user_id) VALUES ($1,$2)",
	// 		c.Descriptions, c.User_id,
	// 	)
	// }
	rows, _ := db.Query("SELECT id,descriptions,created_at,user_id FROM comment")

	for rows.Next() {
		var c Comment
		rows.Scan(
			&c.Id,
			&c.Descriptions,
			&c.Created_at,
			&c.User_id,
		)
		c.TimeFormat = c.Created_at.Format(timeLayout)
		comments = append(comments, c)
	}
	js, _ := json.Marshal(comments)
	os.WriteFile("b.json", js, 0611)

}

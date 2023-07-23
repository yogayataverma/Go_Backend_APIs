package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type tsk struct {
	ID          int    `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Due_Date    string `json:"Due_Date"`
	Status      string `json:"Status"`
}

type tskextraid struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Due_Date    string `json:"Due_Date"`
	ID          int    `json:"ID"`
}

type tskid struct {
	ID int `json:"ID"`
}

func getTask(cs *gin.Context) {

	rows, err := DB.Query("SELECT ID, Title, Description, Due_Date, Status from task")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	tk := make([]tsk, 0)

	for rows.Next() {
		tks := tsk{}
		err = rows.Scan(&tks.ID, &tks.Title, &tks.Description, &tks.Due_Date, &tks.Status)

		if err != nil {
			fmt.Println(err)
		}

		tk = append(tk, tks)
	}

	err = rows.Err()

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if tk == nil {
		cs.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		cs.JSON(http.StatusOK, gin.H{"data": tk})
	}
}

func getTaskById(cs *gin.Context) {

	id := cs.Param("id")

	stmt, err := DB.Prepare("SELECT ID, Title, Description, Due_Date, Status from task WHERE ID = ?")

	if err != nil {
		fmt.Println(err)
	}

	task12 := tsk{}

	sqlErr := stmt.QueryRow(id).Scan(&task12.ID, &task12.Title, &task12.Description, &task12.Due_Date, &task12.Status)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			fmt.Println(sqlErr)
		}
		fmt.Println(task12)
	}
	cs.JSON(http.StatusOK, gin.H{"data": task12})
}

func addTask(cs *gin.Context) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Title: ")
	Title, _ := reader.ReadString('\n')
	fmt.Print("Enter a Description: ")
	Description, _ := reader.ReadString('\n')
	fmt.Print("Enter an Due_Date: ")
	Due_Date, _ := reader.ReadString('\n')

	newtsk := tsk{
		Title:       Title,
		Description: Description,
		Due_Date:    Due_Date,
	}

	tx, err := DB.Begin()

	if err != nil {
		fmt.Println(err)
	}

	stmt, err := tx.Prepare("INSERT INTO task ( Title, Description, Due_Date ) VALUES ( ?, ?, ?)")

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(newtsk.Title, newtsk.Description, newtsk.Due_Date)

	if err != nil {
		fmt.Println(err)
	}

	tx.Commit()

	rows, err := DB.Query("SELECT ID from task ORDER BY ID DESC LIMIT 1")

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	tk := make([]tskid, 0)

	for rows.Next() {
		tks := tskid{}
		err = rows.Scan(&tks.ID)

		if err != nil {
			fmt.Println(err)
		}

		tk = append(tk, tks)
	}

	err = rows.Err()

	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	if tk == nil {
		cs.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		cs.JSON(http.StatusOK, gin.H{"data ID alloted": tk})
	}
}

func updateTask(cs *gin.Context) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a Title: ")
	Title, _ := reader.ReadString('\n')
	fmt.Print("Enter a Description: ")
	Description, _ := reader.ReadString('\n')
	fmt.Print("Enter an Due_Date: ")
	Due_Date, _ := reader.ReadString('\n')

	personId, err1 := strconv.Atoi(cs.Param("id"))

	if err1 != nil {
		fmt.Println(err1)
	}

	newtsk := tskextraid{
		Title:       Title,
		Description: Description,
		Due_Date:    Due_Date,
		ID:          personId,
	}

	tx, err := DB.Begin()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := tx.Prepare("UPDATE task SET Title = ? , Description = ? ,Due_Date = ? WHERE ID = ?")

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(newtsk.Title, newtsk.Description, newtsk.Due_Date, personId)

	if err != nil {
		cs.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
	} else {
		cs.JSON(http.StatusOK, gin.H{"message": "Updated"})
	}

	tx.Commit()

}

func deleteTask(cs *gin.Context) {

	id := cs.Param("id")

	tx, err := DB.Begin()

	if err != nil {
		fmt.Println(err)
	}

	stmt, err := DB.Prepare("DELETE from task where ID = ?")

	if err != nil {
		fmt.Println(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		fmt.Println(err)
	}

	tx.Commit()

	cs.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func options(cs *gin.Context) {
	cs.JSON(http.StatusOK, gin.H{"message": "options Called"})
}

func main() {

	r1 := gin.Default()

	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		fmt.Println(err)
	}

	DB = db

	{
		r1.GET("tasks", getTask)
		r1.GET("tasks/:id", getTaskById)
		r1.POST("tasks", addTask)
		r1.PUT("tasks/:id", updateTask)
		r1.DELETE("tasks/:id", deleteTask)
		r1.OPTIONS("task6", options)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r1.Run(":8080")

}

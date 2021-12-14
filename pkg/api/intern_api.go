package api

import (
	"fmt"
	"log"
	"strconv"

	cn "learn-golang/pkg/connect"
	"learn-golang/pkg/dto"

	"github.com/gin-gonic/gin"
)

func GetAllIntern(c *gin.Context) {
	db := cn.ConnectDB()
	rows, err := db.Query("SELECT * FROM INTERN")

	if err != nil {
		c.JSON(500, gin.H{
			"messages": "No record",
		})
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var interns []dto.Intern

	for rows.Next() {

		var intern dto.Intern
		if err := rows.Scan(&intern.InternId, &intern.InternName, &intern.InternBirthday, &intern.InternInCompany, &intern.Rule); err != nil {
			log.Fatal(err)
			return
		}
		interns = append(interns, intern)
	}

	c.JSON(200, interns)
	defer db.Close()
}

func GetInternById(c *gin.Context) {
	db := cn.ConnectDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "chưa có id",
		})
		return
	}

	var intern dto.Intern

	if err := db.QueryRow("SELECT * from INTERN where INTERN_ID = :1", id).Scan(&intern.InternId, &intern.InternName, &intern.InternBirthday, &intern.InternInCompany, &intern.Rule); err != nil {
		c.JSON(500, gin.H{
			"error": "id chưa tồn tại",
		})
		return
	}

	c.JSON(200, intern)
	defer db.Close()
}

func CreateIntern(c *gin.Context) {
	db := cn.ConnectDB()

	var newintern dto.Intern

	if err := c.ShouldBindJSON(&newintern); err != nil {
		c.JSON(500, gin.H{
			"error": "chưa có body",
		})
		return
	}
	createstaff, err := db.Exec("INSERT INTO INTERN ( INTERN_ID,INTERN_NAME,INTERN_BIRTHDAY,INTERN_IN_COMPANY,RULE) VALUES (:1, :2,TO_DATE(:3, 'yyyy/mm/dd'),TO_DATE(:4, 'yyyy/mm/dd'),:5)", newintern.InternId, newintern.InternName, newintern.InternBirthday, newintern.InternInCompany, newintern.Rule)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "create fail",
		})
		return
	}
	c.JSON(200, createstaff)
	defer db.Close()

}

func UpdateIntern(c *gin.Context) {
	db := cn.ConnectDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "chưa có id",
		})
		return
	}

	var updateIntern dto.Intern

	if err := c.ShouldBindJSON(&updateIntern); err != nil {
		c.JSON(500, gin.H{
			"error": "chưa có body",
		})
		return
	}

	updateIntern.InternId = id

	intern, err := db.Exec(`Update INTERN set INTERN_NAME = :1,INTERN_BIRTHDAY=TO_DATE(:2, 'yyyy/mm/dd'),INTERN_IN_COMPANY=TO_DATE(:3, 'yyyy/mm/dd') where INTERN_ID = :4`, updateIntern.InternName, updateIntern.InternBirthday, updateIntern.InternInCompany, updateIntern.InternId)
	if err != nil {
		c.JSON(500, gin.H{
			"messages": "update fail",
		})
		return
	}

	c.JSON(200, intern)
	defer db.Close()
}

func DeleteIntern(c *gin.Context) {
	db := cn.ConnectDB()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "chưa có id",
		})
		return
	}

	_, error := db.Exec("Delete from INTERN where INTERN_ID = :1", id)
	if error != nil {
		c.JSON(500, gin.H{
			"error": "delete fail",
		})
		return
	}
	c.JSON(200, "Delete success")
	defer db.Close()
}

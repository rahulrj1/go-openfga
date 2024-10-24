package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/test", testHandler)
	router.POST("/import", importHandler)
	router.GET("/export", exportHandler)

	router.Run(":8081")
}

func importHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Create a temporary file to save the uploaded file
	tempFile, err := os.CreateTemp("", "uploaded-file.xlsx")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create temporary file"})
		return
	}
	defer tempFile.Close()

	// Copy the uploaded file to the temporary file
	_, err = io.Copy(tempFile, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file"})
		return
	}

	// Parse the Excel file
	xlFile, err := xlsx.OpenFile(tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open Excel file"})
		return
	}

	// Extract content from the Excel file
	var content string
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				content += cell.String() + "\t"
			}
			content += "\n"
		}
	}

	// Send the stringified content as a response
	c.String(http.StatusOK, content)

	// c.JSON(200, gin.H{"message": "This is the export route of the REST API"})
}

func exportHandler(c *gin.Context) {

	// dummy data
	data := [][]interface{}{
		{"Name", "Age", "City"},
		{"John", 30, "New York"},
		{"Alice", 25, "Los Angeles"},
		{"Bob", 35, "Chicago"},
	}

	// Create a new Excel file
	file := excelize.NewFile()

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new sheet.
	index, err := file.NewSheet("Users")
	if err != nil {
		fmt.Println(err)
		return
	}
	file.SetActiveSheet(index)

	// Add data to the Users sheet
	sheet := "Users"
	for i, row := range data {
		for j, value := range row {
			cell := string(rune('A'+j)) + strconv.Itoa(i+1) // Convert column index to letter
			file.SetCellValue(sheet, cell, value)
		}
	}

	// // Save the Excel file to a temporary location
	// tempFile := "exported_data.xlsx"
	// if err := file.SaveAs(tempFile); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Excel file"})
	// 	return
	// }

	// // Set response headers to force download
	// c.Writer.Header().Set("Content-Disposition", "attachment; filename="+tempFile)
	// c.Writer.Header().Set("Content-Type", "application/octet-stream")
	// c.File(tempFile)

	// Save the Excel file to a buffer
	buffer := new(bytes.Buffer)
	if err := file.Write(buffer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write Excel file"})
		return
	}

	// Set response headers to force download
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=exported_data.xlsx")
	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

	// c.JSON(200, gin.H{"message": "This is the import route of the REST API"})
}

func testHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "This is the test route of the REST API"})
}

package middleware

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"path/filepath"

	"github.com/kataras/iris"
)

// Queue struct with Domain, weight and priority
type Queue struct {
	Domain   string
	Weight   int
	Priority int
}

// Que declaration
var Que []string

var lowQue []string
var medQue []string
var highQue []string

// Repository should implement common methods
type Repository interface {
	Read() []*Queue
}

// GetLastCharAsInt returns the last character of a string as int
func GetLastCharAsInt(s string) int {
	last := string(s[len(s)-1])
	value, err := strconv.Atoi(last)
	if err != nil {
		fmt.Println(value)
	}
	return value
}

func (q *Queue) Read() []*Queue {
	path, _ := filepath.Abs("")
	file, err := os.Open(path + "/api/middleware/domain.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var queueArray []*Queue
	index := 0
	tempQueue := &Queue{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			index = 0
			tempQueue = &Queue{}
			continue
		}
		if index == 0 {
			tempQueue.Domain = text
		} else if index == 1 {
			tempQueue.Weight = GetLastCharAsInt(text)
		} else if index == 2 {
			tempQueue.Priority = GetLastCharAsInt(text)
			queueArray = append(queueArray, tempQueue)
		}
		index++
	}

	return queueArray
}

// PrioritizationValue returns 'low', 'medium', 'high' depending on certain criteria
func PrioritizationValue(weight int, priority int) string {
	result := "low"
	weightValue := "low"
	priorityValue := "low"

	if weight >= 5 {
		weightValue = "medium"
	}
	if priority >= 5 {
		priorityValue = "medium"
	}

	if priorityValue == "medium" && weightValue == "medium" {
		result = "high"
	} else if priorityValue == "medium" || weightValue == "medium" {
		result = "medium"
	}

	return result
}

// ProxyMiddleware should queue our incoming requests
func ProxyMiddleware(c iris.Context) {
	domain := c.GetHeader("domain")
	if len(domain) == 0 {
		c.JSON(iris.Map{"status": 400, "result": "error"})
		return
	}
	var repo Repository
	repo = &Queue{}

	for _, row := range repo.Read() {
		if domain == row.Domain {
			priorityValue := PrioritizationValue(row.Weight, row.Priority)
			if priorityValue == "high" {
				highQue = append(highQue, domain)
			} else if priorityValue == "medium" {
				medQue = append(medQue, domain)
			} else {
				lowQue = append(lowQue, domain)
			}
		}
	}

	highMed := append(highQue, medQue...)
	highMedLow := append(highMed, lowQue...)

	Que = highMedLow

	c.Next()
}

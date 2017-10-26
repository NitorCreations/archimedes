package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
	"encoding/json"

	"github.com/icrowley/fake"
)

var (
	Build string

	Id int = 30

	Priorities = []string{"Low", "Lowest", "Medium", "High", "Highest"}
	Statuses   = []string{"Backlog", "Selected for development", "In progress", "Done"}
	Types      = []string{"Epic", "Story", "Task"}
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Key    string  `json:"key"`
	Issues []Issue `json:"issues"`
}

type CustomField struct {
	Name  string `json:"fieldName"`
	Type  string `json:"fieldType"`
	Value string `json:"value"`
}

type Issue struct {
	Type   string `json:"issueType"`
	Key    string `json:"key"`
	Parent string `json:"parent"`

	Summary     string `json:"summary"`
	Description string `json:"description"`

	CustomFields []CustomField `json:"customFieldValues"`

	Priority string `json:"priority"`
	Status   string `json:"status"`

	Created  time.Time `json:"created"`
	Modified time.Time `json:"updated"`
}

func story(epic string) (story Issue) {
	epiclink := CustomField{"Epic Link", "com.pyxis.greenhopper.jira:gh-epic-link", epic}

	story = Issue{
		"Story",
		"MOCK-" + nextId(),
		"",
		"As " + fake.JobTitle() + " I want " + fake.Word() + " so that " + fake.WordsN(2),
		fake.SentencesN(5),
		[]CustomField{epiclink},
		randomPriority(),
		randomStatus(),
		randomTime(),
		time.Now(),
	}
	return
}

func epic() (epic Issue, epicname CustomField) {
	id := nextId()
	epicname = CustomField{"Epic Name", "com.pyxis.greenhopper.jira:gh-epic-label", fake.Brand() + "-" + id}

	epic = Issue{
		Types[0],
		"MOCK-" + id,
		"",
		"For " + fake.Company() + " who " + fake.Product() + " the " + fake.ProductName() + " is a " + fake.Brand(),
		fake.SentencesN(5),
		[]CustomField{epicname},
		randomPriority(),
		randomStatus(),
		randomTime(),
		time.Now(),
	}
	return
}

func randomTime() (time.Time) {
	days := rand.Intn(100)*24
	hours := rand.Intn(100)

	return time.Now().Add(time.Hour * time.Duration(-days) ).Add(time.Hour * time.Duration(-hours))
}

func randomPriority() (string) {
	return Priorities[rand.Int()%len(Priorities)]
}

func randomStatus() (string) {
	return Statuses[rand.Int()%len(Statuses)]
}

func nextId() (string) {
	Id++
	return strconv.Itoa(Id)
}

func main() {
	rand.Seed(time.Now().Unix())

	var issues []Issue

	for i := 0; i < 10; i++ {
		epic, epicname := epic()
		issues = append(issues, epic)
		for i := 0; i < 10; i++ {
			issues = append(issues, story(epicname.Value))
		}
	}

	project := Project{
		"MOCK",
		issues,
	}
	projects := Projects{[]Project{project}}

	b, _ := json.Marshal(projects)

	fmt.Println(string(b))
}

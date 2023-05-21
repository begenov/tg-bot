package models

import "time"

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatID int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type User struct {
	ID     int
	ChatID int
	Stage  int
	Lang   string
	Name   string
	Phone  string
	Aim    int
	Age    int
	Gender int

	///
	Field   string
	FieldId int
	Job     string
	JobId   int
	Salary  string
}

type JobSeeker struct {
	ID         int
	ChatID     int
	Sphere     string
	Profession string
	Salary     string
}

type Vacancy struct {
	ID               int
	Company          string
	BIN              string
	Sphere           string
	Position         string
	Salary           string
	Requirements     string
	Responsibilities string
	Status           string
	CreationDate     time.Time
	PublicationDate  time.Time
	ModerationDate   time.Time
}

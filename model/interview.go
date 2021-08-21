package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Interview struct {
	UserId         int       `json:"user_id"`
	OrganizationId int       `json:"organization_id"`
	DepartmentName string    `json:"department_name"`
	First          string    `json:"first"`
	FirstStatus    string    `json:"first_status"`
	Second         string    `json:"second"`
	SecondStatus   string    `json:"second_status"`
	Questions      paramList `json:"questions"`
	Scores         paramList `json:"scores"`
}

func (p paramList) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *paramList) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}

type paramList []string

package dto

type Intern struct {
	InternId        int    `json:"INTERN_ID"`
	InternName      string `json:"INTERN_NAME"`
	InternBirthday  string `json:"INTERN_BIRTHDAY"`
	InternInCompany string `json:"INTERN_IN_COMPANY"`
	Rule            string `json:"RULE"`
}

package models

type Ticket struct {
	Id       int    `json:"id"`
	FilePath string `json:"file_path"`
	Notes    string `json:"notes"`
}

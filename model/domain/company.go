package domain

// Company holds information's company
type Company struct {
	Base
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

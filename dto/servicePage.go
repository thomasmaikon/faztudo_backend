package dto

type ServicePageInput struct {
	Id          int
	Name        string
	Image       string
	Value       float64
	Description string
}

type ServicePageOutput struct {
	Id                  int
	Name                string
	Description         string
	Image               string
	Value               float64
	PositiveEvaluations *int
	NegativeEvaluations *int
}

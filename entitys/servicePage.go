package entitys

type ServicePage struct {
	Id                  uint64
	Name                string
	Description         string
	Image               string
	Value               float64
	positiveEvaluations uint64
	negativeEvaluations uint64
}

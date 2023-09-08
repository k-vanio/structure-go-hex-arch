package api

import "github.com/k-vanio/structure-go-hex-arch/internal/ports"

type Adapter struct {
	arith ports.Arithmetic
	db    ports.Db
}

func NewAdapter(arith ports.Arithmetic, db ports.Db) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (api *Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddHistory(answer, "Addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api *Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := api.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddHistory(answer, "Subtraction")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api *Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := api.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddHistory(answer, "Multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api *Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := api.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddHistory(answer, "Division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}

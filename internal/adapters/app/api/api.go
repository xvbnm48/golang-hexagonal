package api

import "github.com/xvbnm48/golang-hexagonal/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (api Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := api.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (api Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := api.arith.Multiplication(a, b)

	if err != nil {
		return 0, nil
	}

	return answer, nil

}

func (api Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := api.arith.Substraction(a, b)
	if err != nil {
		return 0, nil
	}

	return answer, nil
}

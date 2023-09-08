package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (ap *Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (ap *Adapter) Subtraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (ap *Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (ap *Adapter) Division(a, b int32) (int32, error) {
	return a / b, nil
}

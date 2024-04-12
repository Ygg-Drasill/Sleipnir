package ast

func NewExpression(firstOperand, operator, secondOperand Attribute) (Expression, error) {
	return Expression{
		firstOperand:  firstOperand,
		secondOperand: secondOperand,
		operator:      operator,
	}, nil
}

func Add(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 + intVal2, nil
}

func Sub(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 - intVal2, nil
}

func Mul(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 * intVal2, nil
}

func Div(val1, val2 Attribute) (Attribute, error) {
	intVal1, err := toInt64(val1)
	if err != nil {
		return nil, err
	}
	intVal2, err := toInt64(val2)
	if err != nil {
		return nil, err
	}
	return intVal1 / intVal2, nil
}

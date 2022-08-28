package tokens

type Operator struct {
	Prec     int
	Assoc    rune
	Operands uint
}

var Operators map[string]Operator = map[string]Operator{
	"+": {
		Prec:     1,
		Assoc:    'l',
		Operands: 2,
	},
	"-": {
		Prec:     1,
		Assoc:    'l',
		Operands: 2,
	},
	"*": {
		Prec:     2,
		Assoc:    'l',
		Operands: 2,
	},
	"/": {
		Prec:     2,
		Assoc:    'l',
		Operands: 2,
	},
	"%": {
		Prec:     2,
		Assoc:    'l',
		Operands: 2,
	},
	"^": {
		Prec:     3,
		Assoc:    'r',
		Operands: 2,
	},
	"!": {
		Prec:     4,
		Assoc:    'r',
		Operands: 1,
	},
}

func NewOperator(op string) Operator {
	return Operators[op]
}

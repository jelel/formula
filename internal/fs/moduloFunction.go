package fs

import "github.com/yidane/formula/opt"

type ModuloFunction struct {
}

func (*ModuloFunction) Name() string {
	panic("implement me")
}

func (*ModuloFunction) Evaluate(context *opt.FormulaContext, args ...*opt.LogicalExpression) (*opt.Argument, error) {
	panic("implement me")
}

func NewModuloFunction() *ModuloFunction {
	return &ModuloFunction{}
}
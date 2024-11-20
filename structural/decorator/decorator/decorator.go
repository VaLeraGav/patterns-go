package decorator

import "strings"

// -------------- BaseDecorator --------------

type BaseDecorator struct {
	component Component
}

func NewBaseDecorator(component Component) *BaseDecorator {
	return &BaseDecorator{component: component}
}

func (d *BaseDecorator) Operation() string {
	return d.component.Operation()
}

// -------------- UpperCaseDecorator --------------

type UpperCaseDecorator struct {
	BaseDecorator
}

func NewUpperCaseDecorator(component Component) *UpperCaseDecorator {
	return &UpperCaseDecorator{BaseDecorator: *NewBaseDecorator(component)}
}

func (d *UpperCaseDecorator) Operation() string {
	return strings.ToUpper(d.component.Operation())
}

// -------------- TrimDecorator --------------

type TrimDecorator struct {
	BaseDecorator
}

func NewTrimDecorator(component Component) *TrimDecorator {
	return &TrimDecorator{BaseDecorator: *NewBaseDecorator(component)}
}

func (d *TrimDecorator) Operation() string {
	return d.trimMessage(d.component.Operation())
}

func (d *TrimDecorator) trimMessage(message string) string {
	return strings.TrimSpace(message)
}

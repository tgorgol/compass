package model

type Label struct {
	Key string
	Value interface{}
}

type LabelableObject string

const (
	RuntimeLabelableObject LabelableObject = "Runtime"
	ApplicationLabelableObject LabelableObject = "Application"
)
package elastic

type Document interface {
	GetID() string
	SetID(id string)
}

const DefaultTimeout int64 = 3000 // default in milliseconds

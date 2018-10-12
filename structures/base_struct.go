package structures

type BaseStruct interface {
	Description()
}

type baseStruct struct {}

func (b *baseStruct) Description() {}
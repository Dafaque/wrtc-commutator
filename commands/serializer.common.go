package commands

type Serializer func(interface{}) ([]byte, error)

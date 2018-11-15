package data

type Parser interface {
	ParseMapToData(input map[string]interface{}) (*Data, error)
}

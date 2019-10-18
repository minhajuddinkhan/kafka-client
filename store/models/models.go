package models

// All returns all models
func All() []interface{} {

	var models []interface{}
	models = append(models, &Broker{})
	models = append(models, &TLS{})

	return models
}

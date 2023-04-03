package bootstrap

import "go_backend_clean_code/mongo"

type Applicaton struct {
	Env *Env
	Mongo mongo.Client
}

func App() Applicaton {
	app := &Applicaton{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Applicaton) CloseMongoDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
package services

import (
	"justicia/api/app/models/oposicion"
	"justicia/api/config/logging"
	"os"
	"time"

	"github.com/sarulabs/di"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
)

// Services contains the definitions of the application services.
var Services = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return logging.Logger, nil
		},
	},
	{
		Name:  "mongo-pool",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return mgo.DialWithTimeout(os.Getenv("MONGO_URL"), 5*time.Second)
		},
		Close: func(obj interface{}) error {
			obj.(*mgo.Session).Close()
			return nil
		},
	},
	{
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return ctn.Get("mongo-pool").(*mgo.Session).Copy(), nil
		},
		Close: func(obj interface{}) error {
			obj.(*mgo.Session).Close()
			return nil
		},
	},
	{
		Name:  "quiz-repository",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &oposicion.QuizRepository{
				Session: ctn.Get("mongo").(*mgo.Session),
			}, nil
		},
	},
	{
		Name:  "quiz-manager",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &oposicion.QuizManager{
				Repo:   ctn.Get("quiz-repository").(*oposicion.QuizRepository),
				Logger: ctn.Get("logger").(*zap.Logger),
			}, nil
		},
	},
}

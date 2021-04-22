package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *MongoDBLanguage) New() (referenceInicialized interface{}, err error) {
	referenceInicialized = &MongoDBLanguage{}
	err = referenceInicialized.(*MongoDBLanguage).Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.(*MongoDBLanguage).Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

package main

import (
	constants "github.com/helmutkemper/kemper.com.br.plugin.dataaccess.constants"
	"github.com/helmutkemper/util"
)

func (e *MongoDBLanguage) New() (referenceInicialized *MongoDBLanguage, err error) {
	referenceInicialized = &MongoDBLanguage{}
	err = referenceInicialized.Connect(constants.KMongoDBConnectionString)
	if err != nil {
		util.TraceToLog()
		return
	}

	err = referenceInicialized.Install()
	if err != nil {
		util.TraceToLog()
		return
	}

	return
}

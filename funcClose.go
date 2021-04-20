package main

func (e *MongoDBLanguage) Close() (err error) {
	err = e.Client.Disconnect(e.Ctx)
	return
}

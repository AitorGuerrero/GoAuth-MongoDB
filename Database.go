package GoAuthMongoDB

import (
	mgo "gopkg.in/mgo.v2"
	"encoding/json"
	"os"
)

func Database() *mgo.Database {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err := decoder.Decode(&cfg)

	session, err := mgo.Dial("mongodb://"+cfg.User+":"+cfg.Password+"@"+cfg.Host+":"+cfg.Port+"/"+cfg.Database)
	if err != nil {
		panic(err)
	}

	return session.DB("go-auth")
}

package user

import (
	"github.com/AitorGuerrero/UserGo/user"
	GoAuthMgo "github.com/AitorGuerrero/GoAuthMongoDB"

	t "testing"
	"gopkg.in/mgo.v2/bson"
)

var db = GoAuthMgo.Database()
var ut = user.UserSourceTests{
	func() user.Source {
		s := NewSource(db)
		return &s
	},
	func() {
		db.C("users").RemoveAll(bson.M{})
	},
}

func TestUserSource(t *t.T) {
	ut.Run(t)
}

package user

import (
	"github.com/AitorGuerrero/UserGo/user"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Source struct {
	c *mgo.Collection
}

func NewSource(db *mgo.Database) Source {
	return Source{db.C("users")}
}

func (s *Source) Add (u *user.User) (err error) {
	mgou := FromDomainUser(*u)
	err = s.c.Insert(mgou)
	if nil != err {
		err = user.DuplicatedIdError{}
	}

	return
}

func (s Source) Get(uid user.Id) (*user.User, error) {
	mngu = User{}
	err := s.c.Find(bson.M{
		"id": uid.Serialize(),
	}).One(&mngu)
	if err != nil && err.Error() == "not found" {
		return nil, user.NotExistentUser{uid}
	}
	u := FromMngUser(mngu)

	return &u, nil
}

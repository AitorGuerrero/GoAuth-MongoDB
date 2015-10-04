package user

import (
	"github.com/AitorGuerrero/UserGo/user"
)

type User struct {
	Id string
	Passkey string
	Token string
	Namespaces []string
}

func FromDomainUser(u user.User) User {
	return User {
		u.Id().Serialize(),
		string(u.Passkey()),
		u.Token.Serialize(),
		[]string{},
	}
}

func FromMngUser(mngu User) user.User {
	u = user.New(user.ParseId(mngu.Id), user.Passkey(mngu.Passkey))

	return u
}

package user

import (
	"github.com/AitorGuerrero/UserGo/user"

	t "testing"
)

type fakeEncryptor struct {}

func (fakeEncryptor) Encrypt(i string) string {
	return i
}

var e fakeEncryptor
var pe user.PasskeyEncryptor
var i = user.ParseId("de1ec52a-21c6-491c-99f9-6f29bc1d565f")
var p user.Passkey
var u user.User
var mngu User

func beforeEach() {
	e = fakeEncryptor{}
	pe = user.PasskeyEncryptor{e}
	p = pe.Encrypt(i, "passkey")
	u = user.New(i, p)
}

func TestShouldCreateAMngUserFromDomainUser (t *t.T) {
	mngu = FromDomainUser(u)
	if mngu.Id != u.Id().Serialize() {
		t.Error("The Id should be the same")
	}
	if mngu.Passkey != string(u.Passkey()) {
		t.Error("The Id should be the same")
	}
	if mngu.Token != u.Token.Serialize() {
		t.Error("The Token should be the same")
	}
}

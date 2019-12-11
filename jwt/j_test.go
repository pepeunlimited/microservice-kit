package jwt

import (
	"github.com/pepeunlimited/microservice-kit/validator"
	"testing"
	"time"
)

var secret []byte = []byte("v3ry-s3cr3t-k3y")

func TestMicrokitjwt_SignIn(t *testing.T) {
	jwt := NewJWT(secret)
	token, err := jwt.SignIn(2*time.Second, "piiparinen", nil, nil, nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	time.Sleep(1*time.Second)
	verify, err := jwt.VerifyCustomClaims(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if validator.IsEmpty(verify.Username) {
		t.FailNow()
	}
	time.Sleep(2*time.Second)
	err = jwt.Verify(token)
	if err == nil {
		t.FailNow()
	}
	if err != ErrExpired {
		t.Log(err)
		t.FailNow()
	}
}
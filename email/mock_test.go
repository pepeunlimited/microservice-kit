package email

import "testing"

func TestMock_Send(t *testing.T) {
	mock := NewEmailMock(1)
	err := mock.Send(Mail{
		From:    From{Email: "info@pepeunlimited.com", Name: "Piia Niemi"},
		To:      []string{"simo.alakotila@gmail.com"},
		Subject: "TestiSimo",
		Body:    "HelloWorld",
	})
	if err == nil {
		t.FailNow()
	}
	err = mock.Send(Mail{
		From:    From{Email: "info@pepeunlimited.com", Name: "Piia Niemi"},
		To:      []string{"simo.alakotila@gmail.com"},
		Subject: "TestiSimo",
		Body:    "HelloWorld",
	})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
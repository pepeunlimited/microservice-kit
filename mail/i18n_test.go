package mail

import (
	"testing"
)

func TestNewI18n(t *testing.T) {
	i18n, err := NewI18n("fi")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	i18n.Add("lang", "Suomi")
	if len(i18n.Translations()) != 1 {
		t.FailNow()
	}

	i18n.SetMessages()
	i18nEn, err := NewI18n("en")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	i18nEn.Add("lang", "English")
	if len(i18n.Translations()) != 1 {
		t.FailNow()
	}
	i18nEn.SetMessages()

	echo := NewI18nEcho([]string{"fi", "en"})

	if echo.Text(i18n.Language(), "lang") != "Suomi" {
		t.FailNow()
	}
	if echo.Text(i18nEn.Language(), "lang") != "English" {
		t.FailNow()
	}
}

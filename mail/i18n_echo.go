package mail

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
)

type I18nEcho interface {
	Text(lang language.Tag, key string) string
}

type i18necho struct {
	matcher language.Matcher
}

func (echo i18necho) Text(lang language.Tag, key string) string {
	matched,_,_ := echo.matcher.Match(lang)
	return message.NewPrinter(matched).Sprintf(key)
}

func NewI18nEcho(languages []string) I18nEcho {
	tags := make([]language.Tag,0)
	for _, v := range languages {
		t, err := language.Parse(v)
		if err != nil {
			log.Print("i18n-echo: failed to parse language: "+err.Error())
		}
		tags = append(tags, t)
	}
	return i18necho{matcher:language.NewMatcher(tags)}
}
package mail

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
)

type I18n interface {
	Add(key string, text string)
	Translations() []Translation
	SetMessages()
	Language() language.Tag
}

type i18n struct {
	lang language.Tag
	translations []Translation
}

func (i18n i18n) Language() language.Tag {
	return i18n.lang
}

func (i18n i18n) SetMessages() {
	for _, v := range i18n.translations {
		err := message.SetString(i18n.lang, v.key, v.text)
		if err != nil {
			log.Panic("mail: failed to set messages err: "+err.Error())
		}
	}
}

func (i18n i18n) Translations() []Translation {
	return i18n.translations
}

func (i18n *i18n) Add(key string, text string) {
	i18n.translations = append(i18n.translations, Translation{text:text, key:key})
}

type Translation struct {
	key string
	text string
}

func NewI18n(lang string) (I18n, error) {
	translations := make([]Translation, 0)
	t, err := language.Parse(lang)
	if err != nil {
		return nil, err
	}
	return &i18n{translations:translations, lang:t}, nil
}
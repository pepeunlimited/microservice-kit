package mail

import (
	"github.com/pepeunlimited/microservice-kit/misc"
	"log"
	"net/url"
)

func CreateBaseURL(baseURL string) *url.URL {
	u := misc.GetEnv(SmtpUrl, baseURL)
	parsed, err := url.Parse(u)
	if err != nil {
		log.Panic("microservice-kit: error during CreateURL: "+err.Error())
	}
	return parsed
}
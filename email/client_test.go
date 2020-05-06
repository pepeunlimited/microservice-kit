package email

import "testing"

func TestClient_SendHotmail(t *testing.T) {
	NewBuilder().
		From(PiiiaCom, "PIIIA.com").
		To([]string{"simo.alakotila@gmail.com"}).
		Subject("Subject!").
		Content("Content!").Build()
}
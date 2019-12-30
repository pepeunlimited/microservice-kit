package mail

import "testing"

func TestClient_SendHotmail(t *testing.T) {
	client := NewBuilder("", "").
		From(PiiiaCom, "PIIIA.com").
		To([]string{"simo.alakotila@gmail.com"}).
		Subject("Subject!").
		Content("Content!").
		Build(Hotmail)
	err := client.Send()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
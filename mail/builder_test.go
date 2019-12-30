package mail

import (
	"testing"
)

func TestBuild_Build(t *testing.T) {
	client := NewBuilder("", "").From(PiiiaCom, "").To([]string{PepeUnlimited}).Subject("Hello World!").Content("Content!").Build(Mock)
	client.Send()
}
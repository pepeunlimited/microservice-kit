package email

import (
	"testing"
)

func TestBuild_Build(t *testing.T) {
	NewBuilder().From(PiiiaCom, "").To([]string{PepeUnlimited}).Subject("Hello World!").Content("Content!").Build()

}
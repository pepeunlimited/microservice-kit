package rpcz

import (
	"log"
	"testing"
)

func TestAddRoles(t *testing.T) {
	role := AddRoles([]string{"simo","piia"})


	log.Print(GetRoles(role))
}

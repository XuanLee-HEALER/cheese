package entity_test

import (
	"cheese/entity"
	"testing"
)

func TestNewRoleBirth(t *testing.T) {
	rb := entity.NewRoleBirth("5月18日")
	m, d := rb.GetRoleBirthDetail()
	if m != 5 || d != 18 {
		t.Error("new birth obj error")
	}
}

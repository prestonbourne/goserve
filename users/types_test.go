package users

import (
	"testing"
)

func TestNewAccount(t *testing.T) {
	user, err := NewUser("Andrew", "Garfield", "best_spider_man", "tobeysux123")

	if err != nil {
		t.Errorf("Did not create user %+v", err)
	}
	t.Logf("%+v\n", user)

}

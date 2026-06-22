package Users

import "testing"

func TestNewUser(t *testing.T) {
	userName := "Nathan"
	userPassword := "Password"
	user, err := NewUser(userName, userPassword)
	if err != nil {
		t.Errorf("NewUser.uName = %q, NewUser.uPass = %q, Error = %v", user.uName, user.uPass, err)
	}
}

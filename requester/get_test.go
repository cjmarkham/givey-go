package requester

import "testing"

func TestGetUser (t *testing.T) {
    err, data := GetUser("cjmarkham")
    if err != nil {
      t.Errorf("GetUser() responded with error %s", err.Error())
    }

    if data.User.Givey_Tag != "cjmarkham" {
      t.Errorf("GetUser() returned invalid response")
    }
}

func TestGetUsers (t *testing.T) {
    err, data := GetUsers("1", "5")
    if err != nil {
      t.Errorf("GetUsers() responded with error %s", err.Error())
    }

    amount := len(data.Users)
    if amount != 5 {
      t.Errorf("GetUser() returned %d, expected %d", amount, 4)
    }
}

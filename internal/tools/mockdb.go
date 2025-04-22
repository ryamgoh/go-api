package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"betty": {
		AuthToken: "246DEF",
		Username:  "betty",
	},
}

var mockCreditDetails = map[string]CreditDetails{
	"alex": {
		Credits:  123.5,
		Username: "alex",
	},
	"betty": {
		Credits:  2463.3,
		Username: "betty",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}

	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCreditDetails(username string) *CreditDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CreditDetails{}

	clientData, ok := mockCreditDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

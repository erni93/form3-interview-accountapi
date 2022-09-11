package main

import (
	client "erni93/form3-interview-accountapi/client"
	model "erni93/form3-interview-accountapi/models"
	"testing"
)

func buildNewAccountGB() model.AccountData {
	country := "GB"
	return model.AccountData{
		Type:           "accounts",
		ID:             "ad27e262-9605-4b4b-a0e5-3003ea9cc4dc",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Attributes: &model.AccountAttributes{
			Country: &country,
			BankID:  "400300",
			Bic:     "NWBKGB22",
			Name:    []string{"user1"},
		},
	}
}

func TestHappyFlow(t *testing.T) {
	client := client.WithDefaultConfig()

	err := client.IsAvailable()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	_, err = client.GetAccounts()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	created, err := client.CreateAccount(buildNewAccountGB())
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	_, err = client.GetAccount(created.ID)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	err = client.DeleteAccount(created.ID, *created.Version)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
}

func TestGetAccounts(t *testing.T) {
	client := client.WithDefaultConfig()

	accounts, err := client.GetAccounts()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	if len(accounts) != 0 {
		t.Error("expected accounts len to be 0")
	}

	created, err := client.CreateAccount(buildNewAccountGB())
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	accounts, err = client.GetAccounts()
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	if len(accounts) != 1 {
		t.Error("expected accounts len to be 1")
	}
	if accounts[0].ID != created.ID {
		t.Errorf("expected accounts[0].ID to be %s, got %s", accounts[0].ID, created.ID)
	}

	err = client.DeleteAccount(created.ID, *created.Version)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
}

func TestCreateAccount(t *testing.T) {
	client := client.WithDefaultConfig()

	created, err := client.CreateAccount(buildNewAccountGB())
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	if created == nil {
		t.Error("expected created to not be nil")
	}
	if created.Version == nil {
		t.Error("expected created.Version to not be nil")
	}

	err = client.DeleteAccount(created.ID, *created.Version)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	created, err = client.CreateAccount(model.AccountData{ID: "bad-ID"})
	if err == nil {
		t.Error("expected err to not be nil")
	}
	if created != nil {
		t.Error("expected created to be nil")
	}
}

func TestGetAccount(t *testing.T) {
	client := client.WithDefaultConfig()

	created, err := client.CreateAccount(buildNewAccountGB())
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	if created == nil {
		t.Error("expected account to not be nil")
	}

	account, err := client.GetAccount(created.ID)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	if account == nil {
		t.Error("expected account to not be nil")
	}

	if created.ID != account.ID {
		t.Errorf("expected created.ID to be %s, got %s", created.ID, account.ID)
	}

	err = client.DeleteAccount(created.ID, *created.Version)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
}

func TestDeleteAccount(t *testing.T) {
	client := client.WithDefaultConfig()

	created, err := client.CreateAccount(buildNewAccountGB())
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}
	if created == nil {
		t.Error("expected account to not be nil")
	}

	err = client.DeleteAccount(created.ID, *created.Version)
	if err != nil {
		t.Errorf("expected err to be nil, got %s", err)
	}

	err = client.DeleteAccount("bad-ID", 0)
	if err == nil {
		t.Error("expected err to not be nil")
	}
}

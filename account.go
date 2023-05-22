package onepassword

import (
	"time"
)

type AccountType string

const (
	AccountTypeIndividual AccountType = "INDIVIDUAL"
	// TODO
)

type AccountState string

const (
	AccountStateActive AccountState = "ACTIVE"
	// TODO
)

type GetAccountOutput struct {
	CreatedAt time.Time    `json:"created_at,omitempty"`
	Domain    string       `json:"domain,omitempty"`
	ID        string       `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	State     AccountState `json:"state,omitempty"`
	Type      AccountType  `json:"type,omitempty"`
}

func (c *Client) GetAccount() (*GetAccountOutput, error) {
	output := new(GetAccountOutput)
	return output, c.runJson(output, "account", "get")
}

type ListAccountsOutput []struct {
	AccountID string `json:"account_uuid,omitempty"`
	Email     string `json:"email,omitempty"`
	URL       string `json:"url,omitempty"`
	UserID    string `json:"user_uuid,omitempty"`
}

func (c *Client) ListAccounts() (*ListAccountsOutput, error) {
	output := new(ListAccountsOutput)
	return output, c.runJson(output, "account", "list")
}

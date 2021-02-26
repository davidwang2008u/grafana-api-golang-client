package gapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const policyRootUrl = "/api/access-control/policies"

type Policy struct {
	OrgID       int64        `json:"orgId"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	UID         string       `json:"uid,omitempty"`
	PolicyId    int64        `json:"id,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type Permission struct {
	Permission string `json:"permission"`
	Scope      string `json:"scope"`
}

// get
func (c *Client) GetPolicy(uid string) (*Policy, error) {
	p := &Policy{}
	err := c.request("GET", buildUrl(uid), nil, nil, p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (c *Client) NewPolicy(policy Policy) (*Policy, error) {
	data, err := json.Marshal(policy)
	if err != nil {
		return nil, err
	}

	created := &Policy{}

	err = c.request("POST", policyRootUrl, nil, bytes.NewBuffer(data), &created)
	if err != nil {
		return nil, err
	}

	return created, err
}

func (c *Client) UpdatePolicy(policy Policy) error {
	data, err := json.Marshal(policy)
	if err != nil {
		return err
	}

	err = c.request("PUT", buildUrl(policy.UID), nil, bytes.NewBuffer(data), nil)

	return err
}

func (c *Client) DeletePolicy(uid string) error {
	return c.request("DELETE", buildUrl(uid), nil, nil, nil)
}

func buildUrl(uid string) string {
	return fmt.Sprintf("%s/%s", policyRootUrl, uid)
}

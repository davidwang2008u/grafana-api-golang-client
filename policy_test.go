package gapi

import (
	"testing"

	"github.com/gobs/pretty"
)

const (
	newPolicyResponse = `
{
    "id": 8,
    "orgId": 1,
    "uid": "vc3SCSsGz",
    "name": "test:policy",
    "description": "Test policy description",
    "permissions": [
        {
            "id": 6,
            "permission": "test:self",
            "scope": "test:self",
            "updated": "2021-02-22T16:16:05.646913+01:00",
            "created": "2021-02-22T16:16:05.646912+01:00"
        }
    ],
    "updated": "2021-02-22T16:16:05.644216+01:00",
    "created": "2021-02-22T16:16:05.644216+01:00"
}
`

	updatedPolicyResponse = `{"message":"Policy updated"}`
	deletePolicyResponse = `{"message":"Policy deleted"}`
)

func TestNewPolicy(t *testing.T) {
	server, client := gapiTestTools(t, 200, newPolicyResponse)
	defer server.Close()

	policyReq := Policy{
		OrgID:       1,
		Name:        "test:policy",
		Description: "test:policy",
		Permissions: []Permission{
			{
				Permission: "test:self",
				Scope:      "test:self",
			},
		},
	}

	resp, err := client.NewPolicy(policyReq)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(pretty.PrettyFormat(resp))

	if resp != "vc3SCSsGz" {
		t.Error("Not correctly parsing returned policy uid.")
	}
}

func TestUpdatePolicy(t *testing.T) {
	server, client := gapiTestTools(t, 200, updatedPolicyResponse)
	defer server.Close()

	policyReq := Policy{
		OrgID:       1,
		Name:        "test:policy",
		Description: "test:policy",
		Permissions: []Permission{
			{
				Permission: "test:self1",
				Scope:      "test:self1",
			},
		},
	}

	err := client.UpdatePolicy("vc3SCSsGz", policyReq)
	if err != nil {
		t.Error(err)
	}
}

func TestDeletePolicy(t *testing.T) {
	server, client := gapiTestTools(t, 200, deletePolicyResponse)
	defer server.Close()

	err := client.DeletePolicy("vc3SCSsGz")
	if err != nil {
		t.Error(err)
	}
}

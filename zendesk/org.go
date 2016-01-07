package zendesk

import (
	"fmt"
	"time"
)

type Organization struct {
	Id                 *int64                 `json:"id,omitempty"`
	Url                *string                `json:"url,omitempty"`
	ExternalId         *string                `json:"external_id,omitempty"`
	Name               *string                `json:"name,omitempty"`
	CreatedAt          *time.Time             `json:"created_at,omitempty"`
	UpdatedAt          *time.Time             `json:"updated_at,omitempty"`
	DomainNames        *[]string              `json:"domain_names,omitempty"`
	Details            *string                `json:"details,omitempty"`
	Notes              *string                `json:"notes,omitempty"`
	GroupId            *int64                 `json:"group_id,omitempty"`
	SharedTickets      *bool                  `json:"shared_tickets,omitempty"`
	SharedComments     *bool                  `json:"shared_comments,omitempty"`
	OrganizationFields map[string]interface{} `json:"organization_fields,omitempty"`
}

func (c *client) OrganizationGet(id int64) (*Organization, error) {
	out := new(APIPayload)
	err := c.get(fmt.Sprintf("/api/v2/organizations/%d.json", id), out)
	return out.Organization, err
}

func (c *client) OrganizationCreate(org *Organization) (*Organization, error) {
	in := &APIPayload{Organization: org}
	out := new(APIPayload)
	err := c.post("/api/v2/organizations.json", in, out)
	return out.Organization, err
}

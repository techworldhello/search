package mocks

import (
	"encoding/json"
	"github.com/techworldhello/search/pkg/schema"
	"log"
)

func GetTicketTestData() (t []schema.Ticket) {
	if err := json.Unmarshal([]byte(tickets), &t); err != nil {
		log.Printf("%+v", err)
	}
	return t
}

const tickets = `[
  {
    "_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
    "url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
    "external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
    "created_at": "2016-04-28T11:19:34 -10:00",
    "type": "incident",
    "subject": "A Catastrophe in Korea (North)",
    "description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
    "priority": "high",
    "status": "pending",
    "submitter_id": 38,
    "assignee_id": 24,
    "organization_id": 116,
    "tags": [
      "Ohio",
      "Pennsylvania",
      "American Samoa",
      "Northern Mariana Islands"
    ],
    "has_incidents": false,
    "due_at": "2016-07-31T02:37:50 -10:00",
    "via": "web"
  },
  {
    "_id": "1a227508-9f39-427c-8f57-1b72f3fab87c",
    "url": "http://initech.zendesk.com/api/v2/tickets/1a227508-9f39-427c-8f57-1b72f3fab87c.json",
    "external_id": "3e5ca820-cd1f-4a02-a18f-11b18e7bb49a",
    "created_at": "2016-04-14T08:32:31 -10:00",
    "type": "incident",
    "subject": "A Catastrophe in Micronesia",
    "description": "Aliquip excepteur fugiat ex minim ea aute eu labore. Sunt eiusmod esse eu non commodo est veniam consequat.",
    "priority": "low",
    "status": "hold",
    "submitter_id": 71,
    "assignee_id": 38,
    "organization_id": 112,
    "tags": [
      "Puerto Rico",
      "Idaho",
      "Oklahoma",
      "Louisiana"
    ],
    "has_incidents": false,
    "due_at": "2016-08-15T05:37:32 -10:00",
    "via": "chat"
  },
  {
    "_id": "2217c7dc-7371-4401-8738-0a8a8aedc08d",
    "url": "http://initech.zendesk.com/api/v2/tickets/2217c7dc-7371-4401-8738-0a8a8aedc08d.json",
    "external_id": "3db2c1e6-559d-4015-b7a4-6248464a6bf0",
    "created_at": "2016-07-16T12:05:12 -10:00",
    "type": "problem",
    "subject": "A Catastrophe in Hungary",
    "description": "Ipsum fugiat voluptate reprehenderit cupidatat aliqua dolore consequat. Consequat ullamco minim laboris veniam ea id laborum et eiusmod excepteur sint laborum dolore qui.",
    "priority": "normal",
    "status": "closed",
    "submitter_id": 9,
    "assignee_id": 65,
    "organization_id": 105,
    "tags": [
      "Massachusetts",
      "New York",
      "Minnesota",
      "New Jersey"
    ],
    "has_incidents": true,
    "due_at": "2016-08-06T04:16:06 -10:00",
    "via": "web"
  }
]`

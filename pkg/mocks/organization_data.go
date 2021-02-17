package mocks

import (
	"encoding/json"
	"github.com/techworldhello/search/pkg/schema"
	"log"
)

func GetOrganizationTestData() (o []schema.Organization) {
	if err := json.Unmarshal([]byte(organizations), &o); err != nil {
		log.Printf("%+v", err)
	}
	return o
}

const organizations = `[
  {
    "_id": 101,
    "url": "http://initech.zendesk.com/api/v2/organizations/101.json",
    "external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
    "name": "Enthaze",
    "domain_names": [
      "kage.com",
      "ecratic.com",
      "endipin.com",
      "zentix.com"
    ],
    "created_at": "2016-05-21T11:10:28 -10:00",
    "details": "MegaCorp",
    "shared_tickets": false,
    "tags": [
      "Fulton",
      "West",
      "Rodriguez",
      "Farley"
    ]
  },
  {
    "_id": 102,
    "url": "http://initech.zendesk.com/api/v2/organizations/102.json",
    "external_id": "7cd6b8d4-2999-4ff2-8cfd-44d05b449226",
    "name": "Nutralab",
    "domain_names": [
      "trollery.com",
      "datagen.com",
      "bluegrain.com",
      "dadabase.com"
    ],
    "created_at": "2016-04-07T08:21:44 -10:00",
    "details": "Non profit",
    "shared_tickets": false,
    "tags": [
      "Cherry",
      "Collier",
      "Fuentes",
      "Trevino"
    ]
  },
  {
    "_id": 103,
    "url": "http://initech.zendesk.com/api/v2/organizations/103.json",
    "external_id": "e73240f3-8ecf-411d-ad0d-80ca8a84053d",
    "name": "Plasmos",
    "domain_names": [
      "comvex.com",
      "automon.com",
      "verbus.com",
      "gogol.com"
    ],
    "created_at": "2016-05-28T04:40:37 -10:00",
    "details": "",
    "shared_tickets": true,
    "tags": [
      "Parrish",
      "Lindsay",
      "Armstrong",
      "Vaughn"
    ]
  }
]`

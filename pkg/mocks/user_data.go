package mocks

import (
	"encoding/json"
	"github.com/techworldhello/search/pkg/schema"
	"log"
)

func GetUserTestData() (u []schema.User) {
	if err := json.Unmarshal([]byte(users), &u); err != nil {
		log.Printf("%+v", err)
	}
	return u
}

const users = `[
  {
    "_id": 1,
    "url": "http://initech.zendesk.com/api/v2/users/1.json",
    "external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
    "name": "Francisca Rasmussen",
    "alias": "Miss Coffey",
    "created_at": "2016-04-15T05:19:46 -10:00",
    "active": true,
    "verified": true,
    "shared": false,
    "locale": "en-AU",
    "timezone": "Sri Lanka",
    "last_login_at": "2013-08-04T01:03:27 -10:00",
    "email": "coffeyrasmussen@flotonic.com",
    "phone": "8335-422-718",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 119,
    "tags": [
      "Springville",
      "Sutton",
      "Hartsville/Hartley",
      "Diaperville"
    ],
    "suspended": true,
    "role": "admin"
  },
  {
    "_id": 2,
    "url": "http://initech.zendesk.com/api/v2/users/2.json",
    "external_id": "c9995ea4-ff72-46e0-ab77-dfe0ae1ef6c2",
    "name": "Cross Barlow",
    "alias": "Miss Joni",
    "created_at": "2016-06-23T10:31:39 -10:00",
    "active": true,
    "verified": true,
    "shared": false,
    "locale": "zh-CN",
    "timezone": "Armenia",
    "last_login_at": "2012-04-12T04:03:28 -10:00",
    "email": "jonibarlow@flotonic.com",
    "phone": "9575-552-585",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 106,
    "tags": [
      "Foxworth",
      "Woodlands",
      "Herlong",
      "Henrietta"
    ],
    "suspended": false,
    "role": "admin"
  },
  {
    "_id": 3,
    "url": "http://initech.zendesk.com/api/v2/users/3.json",
    "external_id": "85c599c1-ebab-474d-a4e6-32f1c06e8730",
    "name": "Ingrid Wagner",
    "alias": "Miss Buck",
    "created_at": "2016-07-28T05:29:25 -10:00",
    "active": false,
    "verified": false,
    "shared": false,
    "locale": "en-AU",
    "timezone": "Trinidad and Tobago",
    "last_login_at": "2013-02-07T05:53:38 -11:00",
    "email": "buckwagner@flotonic.com",
    "phone": "9365-482-943",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 104,
    "tags": [
      "Mulino",
      "Kenwood",
      "Wescosville",
      "Loyalhanna"
    ],
    "suspended": false,
    "role": "end-user"
  },
  {
    "_id": 4,
    "url": "http://initech.zendesk.com/api/v2/users/4.json",
    "external_id": "37c9aef5-cf01-4b07-af24-c6c49ac1d1c7",
    "name": "Rose Newton",
    "alias": "Mr Cardenas",
    "created_at": "2016-02-09T07:52:10 -11:00",
    "active": true,
    "verified": true,
    "shared": true,
    "locale": "de-CH",
    "timezone": "Netherlands",
    "last_login_at": "2012-09-25T01:32:46 -10:00",
    "email": "cardenasnewton@flotonic.com",
    "phone": "8685-482-450",
    "signature": "Don't Worry Be Happy!",
    "organization_id": 122,
    "tags": [
      "Gallina",
      "Glenshaw",
      "Rowe",
      "Babb"
    ],
    "suspended": true,
    "role": "end-user"
  }
]`

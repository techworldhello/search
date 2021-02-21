package format

import (
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	f := NewSearchFields()

	reflect.DeepEqual(expectedList, f.List())
}

const expectedList = `

****************
Search User with
----------------
_id
url
external_id
name
alias
created_at
active
verified
shared
locale
timezone
last_login_at
email
phone
signature
organization_id
tags
suspended
role

******************
Search Ticket with
------------------
_id
url
external_id
created_at
type
subject
description
priority
status
submitter_id
assignee_id
organization_id
tags
has_incidents
due_at
via

************************
Search Organization with
------------------------
_id
url
external_id
name
domain_names
created_at
details
shared_tickets
tags`

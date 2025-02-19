// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A query filter used by ListUsers and ListGroups. This filter object provides the
// attribute name and attribute value to search users or groups.
type Filter struct {

	// The attribute path that is used to specify which attribute name to search.
	// Length limit is 255 characters. For example, UserName is a valid attribute path
	// for the ListUsers API, and DisplayName is a valid attribute path for the
	// ListGroups API.
	//
	// This member is required.
	AttributePath *string

	// Represents the data for an attribute. Each attribute value is described as a
	// name-value pair.
	//
	// This member is required.
	AttributeValue *string

	noSmithyDocumentSerde
}

// A group object, which contains a specified group’s metadata and attributes.
type Group struct {

	// Contains the group’s display name value. The length limit is 1,024 characters.
	// This value can consist of letters, accented characters, symbols, numbers,
	// punctuation, tab, new line, carriage return, space, and nonbreaking space in
	// this attribute. The characters <>;:% are excluded. This value is specified at
	// the time the group is created and stored as an attribute of the group object in
	// the identity store.
	//
	// This member is required.
	DisplayName *string

	// The identifier for a group in the identity store.
	//
	// This member is required.
	GroupId *string

	noSmithyDocumentSerde
}

// A user object, which contains a specified user’s metadata and attributes.
type User struct {

	// The identifier for a user in the identity store.
	//
	// This member is required.
	UserId *string

	// Contains the user’s user name value. The length limit is 128 characters. This
	// value can consist of letters, accented characters, symbols, numbers, and
	// punctuation. The characters <>;:% are excluded. This value is specified at the
	// time the user is created and stored as an attribute of the user object in the
	// identity store.
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

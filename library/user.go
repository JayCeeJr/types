// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// User is the library representation of a user.
type User struct {
	ID     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Token  *string `json:"token,omitempty"`
	Hash   *string `json:"-"`
	Active *bool   `json:"active,omitempty"`
	Admin  *bool   `json:"admin,omitempty"`
}

// GetID returns the ID field.
//
// When the provided User type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (u *User) GetID() int64 {
	// return zero value if User type or ID field is nil
	if u == nil || u.ID == nil {
		return 0
	}
	return *u.ID
}

// GetName returns the Name field.
//
// When the provided User type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (u *User) GetName() string {
	// return zero value if User type or Name field is nil
	if u == nil || u.Name == nil {
		return ""
	}
	return *u.Name
}

// GetToken returns the Token field.
//
// When the provided User type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (u *User) GetToken() string {
	// return zero value if User type or Token field is nil
	if u == nil || u.Token == nil {
		return ""
	}
	return *u.Token
}

// GetHash returns the Hash field.
//
// When the provided User type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (u *User) GetHash() string {
	// return zero value if User type or Hash field is nil
	if u == nil || u.Hash == nil {
		return ""
	}
	return *u.Hash
}

// GetActive returns the Active field.
//
// When the provided User type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (u *User) GetActive() bool {
	// return zero value if User type or Active field is nil
	if u == nil || u.Active == nil {
		return false
	}
	return *u.Active
}

// GetAdmin returns the Admin field.
//
// When the provided User type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (u *User) GetAdmin() bool {
	// return zero value if User type or Admin field is nil
	if u == nil || u.Admin == nil {
		return false
	}
	return *u.Admin
}

// SetID sets the ID field.
func (u *User) SetID(v int64) {
	if u == nil {
		return
	}
	u.ID = &v
}

// SetName sets the Name field.
func (u *User) SetName(v string) {
	if u == nil {
		return
	}
	u.Name = &v
}

// SetToken sets the Token field.
func (u *User) SetToken(v string) {
	if u == nil {
		return
	}
	u.Token = &v
}

// SetHash sets the Hash field.
func (u *User) SetHash(v string) {
	if u == nil {
		return
	}
	u.Hash = &v
}

// SetActive sets the Active field.
func (u *User) SetActive(v bool) {
	if u == nil {
		return
	}
	u.Active = &v
}

// SetAdmin sets the Admin field.
func (u *User) SetAdmin(v bool) {
	if u == nil {
		return
	}
	u.Admin = &v
}

// String implements the Stringer interface for the User type.
func (u *User) String() string {
	return fmt.Sprintf("%+v", *u)
}

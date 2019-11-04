// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import "fmt"

// Log is the library representation of a log for a step in a build.
type Log struct {
	ID        *int64  `json:"id,omitempty"`
	BuildID   *int64  `json:"build_id,omitempty"`
	RepoID    *int64  `json:"repo_id,omitempty"`
	ServiceID *int64  `json:"service_id,omitempty"`
	StepID    *int64  `json:"step_id,omitempty"`
	Data      *[]byte `json:"data,omitempty"`
}

// GetID returns the ID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetID() int64 {
	// return zero value if Log type or ID field is nil
	if l == nil || l.ID == nil {
		return 0
	}
	return *l.ID
}

// GetBuildID returns the BuildID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetBuildID() int64 {
	// return zero value if Log type or BuildID field is nil
	if l == nil || l.BuildID == nil {
		return 0
	}
	return *l.BuildID
}

// GetRepoID returns the RepoID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetRepoID() int64 {
	// return zero value if Log type or RepoID field is nil
	if l == nil || l.RepoID == nil {
		return 0
	}
	return *l.RepoID
}

// GetServiceID returns the ServiceID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetServiceID() int64 {
	// return zero value if Log type or ServiceID field is nil
	if l == nil || l.ServiceID == nil {
		return 0
	}
	return *l.ServiceID
}

// GetStepID returns the StepID field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetStepID() int64 {
	// return zero value if Log type or StepID field is nil
	if l == nil || l.StepID == nil {
		return 0
	}
	return *l.StepID
}

// GetData returns the Data field.
//
// When the provided Log type is nil, or the field within
// the type is nil, it returns the zero value for the field.
func (l *Log) GetData() []byte {
	// return zero value if Log type or Data field is nil
	if l == nil || l.Data == nil {
		return []byte{}
	}
	return *l.Data
}

// SetID sets the ID field.
func (l *Log) SetID(v int64) {
	if l == nil {
		return
	}
	l.ID = &v
}

// SetStepID sets the StepID field.
func (l *Log) SetStepID(v int64) {
	if l == nil {
		return
	}
	l.StepID = &v
}

// SetBuildID sets the BuildID field.
func (l *Log) SetBuildID(v int64) {
	if l == nil {
		return
	}
	l.BuildID = &v
}

// SetRepoID sets the RepoID field.
func (l *Log) SetRepoID(v int64) {
	if l == nil {
		return
	}
	l.RepoID = &v
}

// SetData sets the Data field.
func (l *Log) SetData(v []byte) {
	if l == nil {
		return
	}
	l.Data = &v
}

// String implements the Stringer interface for the Log type.
func (l *Log) String() string {
	return fmt.Sprintf("%+v", *l)
}

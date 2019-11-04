// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package library

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStep_Getters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	s := &Step{
		ID:           &num64,
		BuildID:      &num64,
		RepoID:       &num64,
		Number:       &num,
		Name:         &str,
		Stage:        &str,
		Status:       &str,
		Error:        &str,
		ExitCode:     &num,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	wantID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantNumber := num
	wantName := str
	wantStage := str
	wantStatus := str
	wantError := str
	wantExitCode := num
	wantCreated := num64
	wantStarted := num64
	wantFinished := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// run test
	gotID := s.GetID()
	gotBuildID := s.GetBuildID()
	gotRepoID := s.GetRepoID()
	gotNumber := s.GetNumber()
	gotName := s.GetName()
	gotStage := s.GetStage()
	gotStatus := s.GetStatus()
	gotError := s.GetError()
	gotExitCode := s.GetExitCode()
	gotCreated := s.GetCreated()
	gotStarted := s.GetStarted()
	gotFinished := s.GetFinished()
	gotHost := s.GetHost()
	gotRuntime := s.GetRuntime()
	gotDistribution := s.GetDistribution()

	if gotID != wantID {
		t.Errorf("GetID is %v, want %v", gotID, wantID)
	}
	if gotBuildID != wantBuildID {
		t.Errorf("GetBuildID is %v, want %v", gotBuildID, wantBuildID)
	}
	if gotRepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", gotRepoID, wantRepoID)
	}
	if gotNumber != wantNumber {
		t.Errorf("GetNumber is %v, want %v", gotNumber, wantNumber)
	}
	if gotName != wantName {
		t.Errorf("GetName is %v, want %v", gotName, wantName)
	}
	if gotStage != wantStage {
		t.Errorf("GetStage is %v, want %v", gotStage, wantStage)
	}
	if gotStatus != wantStatus {
		t.Errorf("GetStatus is %v, want %v", gotStatus, wantStatus)
	}
	if gotError != wantError {
		t.Errorf("GetError is %v, want %v", gotError, wantError)
	}
	if gotExitCode != wantExitCode {
		t.Errorf("GetExitCode is %v, want %v", gotExitCode, wantExitCode)
	}
	if gotCreated != wantCreated {
		t.Errorf("GetCreated is %v, want %v", gotCreated, wantCreated)
	}
	if gotStarted != wantStarted {
		t.Errorf("GetStarted is %v, want %v", gotStarted, wantStarted)
	}
	if gotFinished != wantFinished {
		t.Errorf("GetFinished is %v, want %v", gotFinished, wantFinished)
	}
	if gotHost != wantHost {
		t.Errorf("GetHost is %v, want %v", gotHost, wantHost)
	}
	if gotRuntime != wantRuntime {
		t.Errorf("GetRuntime is %v, want %v", gotRuntime, wantRuntime)
	}
	if gotDistribution != wantDistribution {
		t.Errorf("GetDistribution is %v, want %v", gotDistribution, wantDistribution)
	}
}

func TestLibrary_Step_Setters(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	s := &Step{}

	wantID := num64
	wantBuildID := num64
	wantRepoID := num64
	wantNumber := num
	wantName := str
	wantStage := str
	wantStatus := str
	wantError := str
	wantExitCode := num
	wantCreated := num64
	wantStarted := num64
	wantFinished := num64
	wantHost := str
	wantRuntime := str
	wantDistribution := str

	// run test
	s.SetID(wantID)
	s.SetBuildID(wantBuildID)
	s.SetRepoID(wantRepoID)
	s.SetNumber(wantNumber)
	s.SetName(wantName)
	s.SetStage(wantStage)
	s.SetStatus(wantStatus)
	s.SetError(wantError)
	s.SetExitCode(wantExitCode)
	s.SetCreated(wantCreated)
	s.SetStarted(wantStarted)
	s.SetFinished(wantFinished)
	s.SetHost(wantHost)
	s.SetRuntime(wantRuntime)
	s.SetDistribution(wantDistribution)

	if *s.ID != wantID {
		t.Errorf("GetID is %v, want %v", *s.ID, wantID)
	}
	if *s.BuildID != wantBuildID {
		t.Errorf("GetBuildID is %v, want %v", *s.BuildID, wantBuildID)
	}
	if *s.RepoID != wantRepoID {
		t.Errorf("GetRepoID is %v, want %v", *s.RepoID, wantRepoID)
	}
	if *s.Number != wantNumber {
		t.Errorf("GetNumber is %v, want %v", *s.Number, wantNumber)
	}
	if *s.Name != wantName {
		t.Errorf("GetName is %v, want %v", *s.Name, wantName)
	}
	if *s.Stage != wantStage {
		t.Errorf("GetStage is %v, want %v", *s.Stage, wantStage)
	}
	if *s.Status != wantStatus {
		t.Errorf("GetStatus is %v, want %v", *s.Status, wantStatus)
	}
	if *s.Error != wantError {
		t.Errorf("GetError is %v, want %v", *s.Error, wantError)
	}
	if *s.ExitCode != wantExitCode {
		t.Errorf("GetExitCode is %v, want %v", *s.ExitCode, wantExitCode)
	}
	if *s.Created != wantCreated {
		t.Errorf("GetCreated is %v, want %v", *s.Created, wantCreated)
	}
	if *s.Started != wantStarted {
		t.Errorf("GetStarted is %v, want %v", *s.Started, wantStarted)
	}
	if *s.Finished != wantFinished {
		t.Errorf("GetFinished is %v, want %v", *s.Finished, wantFinished)
	}
	if *s.Host != wantHost {
		t.Errorf("GetHost is %v, want %v", *s.Host, wantHost)
	}
	if *s.Runtime != wantRuntime {
		t.Errorf("GetRuntime is %v, want %v", *s.Runtime, wantRuntime)
	}
	if *s.Distribution != wantDistribution {
		t.Errorf("GetDistribution is %v, want %v", *s.Distribution, wantDistribution)
	}
}

func TestStep_Getters_Empty(t *testing.T) {
	// setup types
	s := &Step{}

	// run test
	gotID := s.GetID()
	gotBuildID := s.GetBuildID()
	gotRepoID := s.GetRepoID()
	gotNumber := s.GetNumber()
	gotName := s.GetName()
	gotStage := s.GetStage()
	gotStatus := s.GetStatus()
	gotError := s.GetError()
	gotExitCode := s.GetExitCode()
	gotCreated := s.GetCreated()
	gotStarted := s.GetStarted()
	gotFinished := s.GetFinished()
	gotHost := s.GetHost()
	gotRuntime := s.GetRuntime()
	gotDistribution := s.GetDistribution()

	if gotID != 0 {
		t.Errorf("GetID is %v, want 0", gotID)
	}
	if gotBuildID != 0 {
		t.Errorf("GetBuildID is %v, want 0", gotBuildID)
	}
	if gotRepoID != 0 {
		t.Errorf("GetRepoID is %v, want 0", gotRepoID)
	}
	if gotNumber != 0 {
		t.Errorf("GetNumber is %v, want 0", gotNumber)
	}
	if gotName != "" {
		t.Errorf("GetName is %v, want \"\"", gotName)
	}
	if gotStage != "" {
		t.Errorf("GetStage is %v, want \"\"", gotStage)
	}
	if gotStatus != "" {
		t.Errorf("GetStatus is %v, want \"\"", gotStatus)
	}
	if gotError != "" {
		t.Errorf("GetError is %v, want \"\"", gotError)
	}
	if gotExitCode != 0 {
		t.Errorf("GetExitCode is %v, want 0", gotExitCode)
	}
	if gotCreated != 0 {
		t.Errorf("GetCreated is %v, want 0", gotCreated)
	}
	if gotStarted != 0 {
		t.Errorf("GetStarted is %v, want 0", gotStarted)
	}
	if gotFinished != 0 {
		t.Errorf("GetFinished is %v, want 0", gotFinished)
	}
	if gotHost != "" {
		t.Errorf("GetHost is %v, want \"\"", gotHost)
	}
	if gotRuntime != "" {
		t.Errorf("GetRuntime is %v, want \"\"", gotRuntime)
	}
	if gotDistribution != "" {
		t.Errorf("GetDistribution is %v, want \"\"", gotDistribution)
	}
}

func TestStep_String(t *testing.T) {
	// setup types
	num := 1
	num64 := int64(num)
	str := "foo"
	s := &Step{
		ID:           &num64,
		BuildID:      &num64,
		RepoID:       &num64,
		Number:       &num,
		Name:         &str,
		Stage:        &str,
		Status:       &str,
		Error:        &str,
		ExitCode:     &num,
		Created:      &num64,
		Started:      &num64,
		Finished:     &num64,
		Host:         &str,
		Runtime:      &str,
		Distribution: &str,
	}
	want := fmt.Sprintf("%+v", *s)

	// run test
	got := s.String()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("String is %v, want %v", got, want)
	}
}

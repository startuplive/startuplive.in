package models

import (
	"testing"

	// "github.com/ungerik/go-start/mongo"
)

var event = Event{
	Status:              "Public",
	Name:                "Test",
	Topic:               "Testing",
	DescriptionLead:     "Short Description",
	Description:         "Description",
	HowItsDone:          "How",
	Prizes:              "Prizes",
	Language:            "en",
	RegistrationButton:  "Registration",
	RegistrationTagline: "Tagline",
}

// func Test_MongoInitLocalhost(t *testing.T) {
// 	//create Database Connection
// 	///////////////////////////////////////////////////////////////////////////
// 	// Load configuration

// 	err := mongo.InitLocalhost("startuplive_testing", "", "")
// 	if err != nil {
// 		t.Error("InitLocalhost failed")
// 	} else {
// 		t.Log("Test passed")
// 	}
// 	dbname := mongo.Config.Database
// 	if dbname == "startuplive_testing" {
// 		t.Log("Test passed")
// 	} else {
// 		t.Error("DB not set in Config")
// 	}
// }

// func Test_CreateEvent(t *testing.T) {
// 	// mongo.InitLocalhost("startuplive_testing", "", "")

// 	evt := event
// 	err := evt.Save()
// 	if err != nil {
// 		t.Error("Saving the Event failed")
// 	} else {
// 		t.Log("Test passed")
// 	}

// }

func Test_IsPublished(t *testing.T) {
	evt := event

	if evt.IsPublished() {
		t.Log("Test passed")
	} else {
		t.Error("IsPublished did not work as expected.") // log error if it did not work as expected
	}

	evt.Status = "Planned"

	if !evt.IsPublished() {
		t.Log("Test passed")
	} else {
		t.Error("IsPublished did not work as expected.") // log error if it did not work as expected
	}
}

func Test_AboutDone(t *testing.T) {
	evt := event

	if evt.AboutDone() {
		t.Log("Test passed")
	} else {
		t.Error("AboutDone did not work as expected.") // log error if it did not work as expected
	}

	evt.Topic = ""

	if !evt.AboutDone() {
		t.Log("Test passed")
	} else {
		t.Error("AboutDone did not work as expected.") // log error if it did not work as expected
	}
}

func Benchmark_AboutDone(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	evt := event
	evt.AboutDone()
}

func Benchmark_TimeConsumingAboutDone(b *testing.B) { //benchmark function starts with "Benchmark" and takes a pointer to type testing.B
	b.StopTimer() //stop the performance timer temporarily while doing initialization

	//do any time consuming initialization functions here ... 
	//database connection, reading files, network connection, etc.
	evt := event

	b.StartTimer() //restart timer

	evt.AboutDone()
}

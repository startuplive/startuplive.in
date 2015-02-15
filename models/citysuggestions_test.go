package models

import (
	"github.com/ungerik/go-start/mongo"
	"strconv"
	"testing"
)

func Test_MongoInitLocalhost(t *testing.T) {
	//create Database Connection
	err := mongo.InitLocalhost("startuplive_testing", "", "")
	if err != nil {
		t.Error("InitLocalhost failed")
	} else {
		t.Log("Test passed")
	}
	dbname := mongo.Config.Database
	if dbname == "startuplive_testing" {
		t.Log("Test passed")
	} else {
		t.Error("DB not set in Config")
	}
}

func Test_SaveAndRemove(t *testing.T) {
	ctype := "Startup Live Event"
	cdate := "2006-01-02 15:04:05"
	cname := "Test City"
	cemail := "test@example.com"

	city := NewCitySuggestion()
	city.Type.Set(ctype)
	city.Date.SetString(cdate)
	city.Name.Set(cname)
	city.Email.Set(cemail)
	city.Save()

	t.Log(city)
	err := city.Save()
	if err != nil {
		t.Error("error on city save: ", err)
	} else {
		t.Log("Saved City passed")
	}

	var c CitySuggestion
	CitySuggestions.DocumentWithID(city.Ref().ID, &c)
	if c.Type.String() != ctype &&
		c.Date.Get() != cdate &&
		c.Name.Get() != cname &&
		c.Email.Get() != cemail {
		t.Error("City not found")
	} else {
		t.Log("city found")
	}

	err = city.Delete()
	if err != nil {
		t.Error("error on city remove: ", err)
	} else {
		t.Log("Removed City passed")
	}
}

func Test_RemoveAll(t *testing.T) {
	ctype := "Startup Live Event"
	cdate := "2006-01-02 15:04:05"
	cname := "Test City"
	cemail := "test@example.com"

	city := NewCitySuggestion()
	city.Type.Set(ctype)
	city.Date.SetString(cdate)
	city.Name.Set(cname)
	city.Email.Set(cemail)
	city.Save()

	t.Logf("city Ref: ", city.Ref())

	var c CitySuggestion
	CitySuggestions.DocumentWithID(city.Ref().ID, &c)
	if c.Type.String() != ctype &&
		c.Date.Get() != cdate &&
		c.Name.Get() != cname &&
		c.Email.Get() != cemail {
		t.Error("City not found")
	} else {
		t.Log("city found")
	}

	i := CitySuggestions.Iterator()

	var ctemp CitySuggestion
	count := 0
	for i.Next(&ctemp) {
		count++
	}
	if count != 1 {
		t.Error("Iterator Error")
	}

	CitySuggestions.RemoveAll()

	i = CitySuggestions.Iterator()

	count = 0
	for i.Next(&ctemp) {
		count++
	}
	if count != 0 {
		t.Error("RemoveAll Error")
	}
}

func Test_GetItems(t *testing.T) {
	CitySuggestions.RemoveAll()

	city := NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-02 15:04:05")
	city.Name.Set("Test City1")
	city.Email.Set("test@example.com")
	city.Save()

	city = NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-02 15:04:05")
	city.Name.Set("Test City2")
	city.Email.Set("test@example.com")
	city.Save()

	city = NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-02 15:04:05")
	city.Name.Set("Test City3")
	city.Email.Set("test@example.com")
	city.Save()

	i := CitySuggestions.Iterator()

	var c CitySuggestion
	count := 0
	for i.Next(&c) {
		count++
	}
	if count != 3 {

	}

	CitySuggestions.RemoveAll()

}

func Test_SortByNameIterator(t *testing.T) {
	CitySuggestions.RemoveAll()

	city := NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-02 15:04:05")
	city.Name.Set("Test City1")
	city.Email.Set("test@example.com")
	city.Save()

	city = NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-03-02 15:04:05")
	city.Name.Set("Test City2")
	city.Email.Set("test@example.com")
	city.Save()

	city = NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-01 15:04:05")
	city.Name.Set("Test City3")
	city.Email.Set("test@example.com")
	city.Save()

	i := CitySuggestions.Sort("Name").Iterator()

	var c CitySuggestion
	count := 1
	for i.Next(&c) {

		if c.Name.Get() != "Test City"+strconv.Itoa(count) {
			t.Error("sort error")
		}
		count++
	}

	CitySuggestions.RemoveAll()

}

func Benchmark_TimeConsumingSortByNameIterator(b *testing.B) {
	b.StopTimer() //stop the performance timer temporarily while doing initialization

	//do any time consuming initialization functions here ... 
	//database connection, reading files, network connection, etc.
	CitySuggestions.RemoveAll()

	city := NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-02 15:04:05")
	city.Name.Set("Test City1")
	city.Email.Set("test@example.com")
	city.Save()

	city = NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-03-02 15:04:05")
	city.Name.Set("Test City2")
	city.Email.Set("test@example.com")
	city.Save()

	city = NewCitySuggestion()
	city.Type.Set("Startup Live Event")
	city.Date.SetString("2006-01-01 15:04:05")
	city.Name.Set("Test City3")
	city.Email.Set("test@example.com")
	city.Save()

	b.StartTimer() //restart timer

	i := CitySuggestions.Sort("Name").Iterator()

	var c CitySuggestion
	count := 1
	for i.Next(&c) {

		if c.Name.Get() != "Test City"+strconv.Itoa(count) {
			b.Error("sort error")
		}
		count++
	}

	CitySuggestions.RemoveAll()
}

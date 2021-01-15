package config

import (
  "testing"
  "strings"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("%s != %s", a, b)
	}
}

// Test with no config file specified
func TestLoadConfigNoFile( t *testing.T){
  var c Config

  err := LoadConfig("", &c)
  expectedError := "No config file specified"
  if (err==nil){
    t.Error("Did not return error when no file provided expected non nil err actual nil err")
  } else if (err.Error() != expectedError){
    t.Errorf("Did not return correct error, expected [%s] actual [%s]", expectedError, err.Error())
  }

}

// Test with a config file which doesn't exist
func TestLoadConfigMissingFile( t *testing.T){
  var c Config

  err := LoadConfig("../samples/no.cf", &c)
  s := "no such file or directory"
  if (err==nil){
    t.Error("Did not return error when no file provided expected non nil err actual nil err")
  } else {
    if !strings.HasSuffix(err.Error(), s) {
      t.Errorf("Did not return correct error, expected suffix [%s] actual [%s]", s, err.Error())
    }
  }
}

// Test with a valid config file
func TestLoadConfig( t *testing.T){
  var c Config
  // Test with missing config file
  err := LoadConfig("../samples/accounts.cf", &c)

  if (err!=nil){
    t.Error("Error loading config file :", err)
  }
  assertEqual(t, c.Type, "accounts")
  assertEqual(t, c.Header, false)
  assertEqual(t, c.Delimiter, "")
  if (len(c.Fields)!=6){
    t.Errorf("Length of Fields array incorrect expected 6 actual %d", len(c.Fields))
  }
  assertEqual(t, c.Fields[5].Name, "customerId")
  assertEqual(t, c.Fields[5].Width, 6)
  assertEqual(t, c.Fields[5].Mode, "token")
  assertEqual(t, c.Fields[5].RegEx, "[A-Z]{3}[0-9]{3}")
  assertEqual(t, c.Fields[5].PresentRatio, float32(1.0))
}

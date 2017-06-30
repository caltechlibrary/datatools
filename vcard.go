package datatools

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	//"time"
)

/*
Example vcard 4.0 data:

BEGIN:VCARD
VERSION:4.0
N:Doe;John;
FN:John Doe
ORG:Example Science Institute;
TITLE:Professor
EMAIL;PREF;INTERNET:john.doe@example.edu
SOURCE;VALUE=uri:http://directory.example.edu/personnel/john.doe.vcf
REV:2017-06-29T15:01:36.657947
END:VCARD

*/

// VCard is a struct based on VCard V4.0 example
type VCard struct {
	XMLName      xml.Name          `json:"-"`
	Version      string            `xml:"version" json:"version"`
	Name         []string          `xml:"name" json:"name"`
	FullName     string            `xml:"full_name" json:"full_name"`
	Organization []string          `xml:"organization" json:"organization,omitempty"`
	Title        string            `xml:"title,omitempty" json:"title,omitempty"`
	EMail        []string          `xml:"email,omitempty" json:"email,omitempty"`
	Source       string            `xml:"source" json:"source"`
	Revision     string            `xml:"revision" json:"revision"`
	Ext          map[string]string `xml:"ext,omitempty" json:"ext,omitempty"`
}

func (vcard *VCard) Parse(src []byte) error {
	var (
		err          error
		field, value string
		inVCard      bool
	)
	// Break out text into lines
	lines := bytes.Split(src, []byte("\n"))
	for i, line := range lines {
		if bytes.Compare(line, []byte("BEGIN:VCARD")) == 0 {
			if inVCard == true {
				err = fmt.Errorf("line %d, unexpected BEGIN:VCARD", i)
				break
			}
			inVCard = true
			field = ""
			value = ""
		}
		if inVCard == true {
			if bytes.Compare(line, []byte("END:VCARD")) == 0 {
				inVCard = false
				break
			} else {
				// Are we starting a field or are we in another field?
				if bytes.Contains(line, []byte(":")) == true {
					parts := bytes.SplitN(line, []byte(":"), 2)
					field = fmt.Sprintf("%s", parts[0])
					value = fmt.Sprintf("%s", parts[1])
					switch field {
					case "BEGIN":
					case "N":
						vcard.Name = strings.Split(value, ";")
					case "FN":
						vcard.FullName = value
					case "ORG":
						vcard.Organization = strings.Split(value, ";")
					case "TITLE":
						vcard.Title = value
					case "EMAIL":
						vcard.EMail = strings.Split(value, ";")
					case "SOURCE":
						vcard.Source = value
					case "REV":
						vcard.Revision = value
					default:
						vcard.Ext[field] = value
					}
				}
			}
		}
	}
	if inVCard == true && err == nil {
		err = fmt.Errorf("line %d unexpected end of vCard", len(lines))
	}
	return err
}

// AsJSON takes a vcard and returns it as a JSON doc
func (vcard *VCard) AsJSON() ([]byte, error) {
	return json.Marshal(vcard)
}

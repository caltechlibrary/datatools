package codemeta

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	// Caltech Library package
	"github.com/caltechlibrary/doitools"
)

type PersonOrOrganization struct {
	// ORCID is use for person id
	Id   string `json:"@id"`
	Type string `json:"@type"`
	// Name is used by organizations
	Name string `json:"name,omitempty"`
	// Given/Family are used by individual persons
	GivenName   string `json:"givenName,omitempty"`
	FamilyName  string `json:"familyName,omitempty"`
	Affiliation string `json:"affiliation,omitempty"`
	Email       string `json:"email,omitempty"`
}

type Codemeta struct {
	// Terms used by Caltech Library codemeta.json
	// Id should be the DOI if available
	Context           string                  `json:"@context"`
	Type              string                  `json:"@type"`
	Name              string                  `json:"name"`
	Description       string                  `json:"description"`
	CodeRepository    string                  `json:"codeRepository"`
	IssueTracker      string                  `json:"issueTracker"`
	License           string                  `json:"license"`
	Version           string                  `json:"version"`
	Author            []*PersonOrOrganization `json:"author"`
	Contributor       []*PersonOrOrganization `json:"contributor,omitempty"`
	Editor            []*PersonOrOrganization `json:"editor,omitempty"`
	DevelopmentStatus string                  `json:"developmentStatus"`
	DownloadURL       string                  `json:"downloadUrl"`
	Keywords          []string                `json:"keywords"`
	Maintainer        string                  `json:"maintainer,omitempty"`
	Funder            []*PersonOrOrganization `json:"funder,omitempty"`
	CopyrightHolder   []*PersonOrOrganization `json:"copyrightHolder,omitempty"`
	CopyrightYear     string                  `json:"copyrightYear,omitempty"`
	Created           string                  `json:"dateCreated,omitempty"`
	Updated           string                  `json:"dateModified,omitempty"`
	Published         string                  `json:"datePublished,omitempty"`
	Identifier        string                  `json:"identifier,omitempty"` //FIXME: Is this where I can put the DOI

	// Additional codemeta Terms are defined at https://codemeta.github.io/terms/
}

func (person *PersonOrOrganization) ToJSON() ([]byte, error) {
	return json.MarshalIndent(person, "", "\t")
}

func (cm *Codemeta) ToJSON() ([]byte, error) {
	return json.MarshalIndent(cm, "", "\t")
}

func (person *PersonOrOrganization) ToCFF() ([]byte, error) {
	if (person.FamilyName == "") || (person.GivenName == "") || (strings.HasPrefix(person.Id, "https://orcid.org/") == false) {
		return []byte(""), fmt.Errorf("Missing family name, given name or ORCID")
	}
	return []byte(fmt.Sprintf(`
   - family-names: %s
     given-names: %s
     orcid: %s`, person.FamilyName, person.GivenName, person.Id)), nil
}

// ToCff crosswalks a Codemeta data structure rendering
// CITATION.cff document as an array of byte.
// Based on documentation at https://citation-file-format.github.io/
func (cm *Codemeta) ToCff() ([]byte, error) {
	src := []byte(`
cff-version: 1.1.0
message: "If you use this software, please cite it as below."
authors:`)
	for _, person := range cm.Author {
		if text, err := person.ToCFF(); err == nil {
			src = append(src, text...)
		}
	}
	if strings.HasPrefix(cm.Identifier, "https://doi.org/") {
		if doi, err := doitools.NormalizeDOI(cm.Identifier); err == nil {
			src = append(src, []byte(fmt.Sprintf(`
doi: %s`, doi))...)
		}
	}
	if cm.Version != "" {
		src = append(src, []byte(fmt.Sprintf(`
version: %s`, cm.Version))...)
		if cm.Name != "" {
			src = append(src, []byte(fmt.Sprintf(`
title: %s`, cm.Name))...)
		}
	}
	if cm.Published != "" {
		src = append(src, []byte(fmt.Sprintf(`
date-released: %s`, cm.Published))...)
	} else {
		now := time.Now()
		dt := now.Format("2006-01-02")
		src = append(src, []byte(fmt.Sprintf(`
date-released: %s`, dt))...)
	}
	// Add a trailing NL
	src = append(src, []byte("\n")...)
	return src, nil
}

// Show returns a the value as a string for the json path described by
// dataPaths. If dataPath is "." return the whole object as a string.
// if the dataPath is not valid return an empty string.
func (cm *Codemeta) Show(dataPath string) (string, error) {
	return "", fmt.Errorf("Show() not implemented.")
}

// Set sets the value in the codemeta.json file for a provided json path.
func (cm *Codemeta) Set(dataPath string, value string) error {
	return fmt.Errorf("Set() not implemented.")
}

// Delete removes the value in the codemeta.json file for a provided json path.
func (cm *Codemeta) Delete(dataPath string) error {
	return fmt.Errorf("Delete() not implemented.")
}

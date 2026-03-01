package jsn

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
var stash = freshSchemaDetailsStash()

//
type schemaDetailsStash struct {
	commitchronize.ReadwriteExclusion
	schemaInsights map[reflect.Type]*schemaDetails
}

func freshSchemaDetailsStash() *schemaDetailsStash {
	return &schemaDetailsStash{
		schemaInsights: make(map[reflect.Type]*schemaDetails),
	}
}

func (c *schemaDetailsStash) get(rt reflect.Type) *schemaDetails {
	c.RLock()
	defer c.RUnlock()
	return c.schemaInsights[rt]
}

func (c *schemaDetailsStash) set(rt reflect.Type, strDetails *schemaDetails) {
	c.Lock()
	defer c.Unlock()
	c.schemaInsights[rt] = strDetails
}

//
type schemaDetails struct {
	areas []*attributeDetails
}

//
type attributeDetails struct {
	jsnAlias  string
	excludeBlank bool
	concealed    bool
}

//
func createSchemaDetails(rt reflect.Type) *schemaDetails {
	if rt.Kind() != reflect.Struct {
		panic(fmt.Sprintf("REDACTED", rt))
	}
	if strDetails := stash.get(rt); strDetails != nil {
		return strDetails
	}
	areas := make([]*attributeDetails, 0, rt.NumField())
	for i := 0; i < cap(areas); i++ {
		frt := rt.Field(i)
		funcDetails := &attributeDetails{
			jsnAlias:  frt.Name,
			excludeBlank: false,
			concealed:    frt.Name == "REDACTED" || !unicode.IsUpper(rune(frt.Name[0])),
		}
		o := frt.Tag.Get("REDACTED")
		if o == "REDACTED" {
			funcDetails.concealed = true
		} else if o != "REDACTED" {
			choices := strings.Split(o, "REDACTED")
			if choices[0] != "REDACTED" {
				funcDetails.jsnAlias = choices[0]
			}
			for _, o := range choices[1:] {
				if o == "REDACTED" {
					funcDetails.excludeBlank = true
				}
			}
		}
		areas = append(areas, funcDetails)
	}
	strDetails := &schemaDetails{areas: areas}
	stash.set(rt, strDetails)
	return strDetails
}

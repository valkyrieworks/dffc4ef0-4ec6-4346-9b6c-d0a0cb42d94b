package json

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	engineconnect "github.com/valkyrieworks/utils/align"
)

//
var repository = newStructDetailsRepository()

//
type structDetailsRepository struct {
	engineconnect.ReadwriteLock
	structDetails map[reflect.Type]*structDetails
}

func newStructDetailsRepository() *structDetailsRepository {
	return &structDetailsRepository{
		structDetails: make(map[reflect.Type]*structDetails),
	}
}

func (c *structDetailsRepository) get(rt reflect.Type) *structDetails {
	c.RLock()
	defer c.RUnlock()
	return c.structDetails[rt]
}

func (c *structDetailsRepository) set(rt reflect.Type, sDetails *structDetails) {
	c.Lock()
	defer c.Unlock()
	c.structDetails[rt] = sDetails
}

//
type structDetails struct {
	attributes []*fieldDetails
}

//
type fieldDetails struct {
	jsonLabel  string
	ignoreEmpty bool
	concealed    bool
}

//
func createStructDetails(rt reflect.Type) *structDetails {
	if rt.Kind() != reflect.Struct {
		panic(fmt.Sprintf("REDACTED", rt))
	}
	if sDetails := repository.get(rt); sDetails != nil {
		return sDetails
	}
	attributes := make([]*fieldDetails, 0, rt.NumField())
	for i := 0; i < cap(attributes); i++ {
		frt := rt.Field(i)
		fDetails := &fieldDetails{
			jsonLabel:  frt.Name,
			ignoreEmpty: false,
			concealed:    frt.Name == "REDACTED" || !unicode.IsUpper(rune(frt.Name[0])),
		}
		o := frt.Tag.Get("REDACTED")
		if o == "REDACTED" {
			fDetails.concealed = true
		} else if o != "REDACTED" {
			opts := strings.Split(o, "REDACTED")
			if opts[0] != "REDACTED" {
				fDetails.jsonLabel = opts[0]
			}
			for _, o := range opts[1:] {
				if o == "REDACTED" {
					fDetails.ignoreEmpty = true
				}
			}
		}
		attributes = append(attributes, fDetails)
	}
	sDetails := &structDetails{attributes: attributes}
	repository.set(rt, sDetails)
	return sDetails
}

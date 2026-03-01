package customer

//
//
type IfaceInquireChoices struct {
	Altitude int64
	Validate  bool
}

//
var FallbackIfaceInquireChoices = IfaceInquireChoices{Altitude: 0, Validate: false}

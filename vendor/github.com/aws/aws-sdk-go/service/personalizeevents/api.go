// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalizeevents

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

const opPutEvents = "PutEvents"

// PutEventsRequest generates a "aws/request.Request" representing the
// client's request for the PutEvents operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutEvents for more information on using the PutEvents
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the PutEventsRequest method.
//    req, resp := client.PutEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutEvents
func (c *PersonalizeEvents) PutEventsRequest(input *PutEventsInput) (req *request.Request, output *PutEventsOutput) {
	op := &request.Operation{
		Name:       opPutEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/events",
	}

	if input == nil {
		input = &PutEventsInput{}
	}

	output = &PutEventsOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PutEvents API operation for Amazon Personalize Events.
//
// Records user interaction event data.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Personalize Events's
// API operation PutEvents for usage and error information.
//
// Returned Error Types:
//   * InvalidInputException
//   Provide a valid value for the field or parameter.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/personalize-events-2018-03-22/PutEvents
func (c *PersonalizeEvents) PutEvents(input *PutEventsInput) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	return out, req.Send()
}

// PutEventsWithContext is the same as PutEvents with the addition of
// the ability to pass a context and additional request options.
//
// See PutEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PersonalizeEvents) PutEventsWithContext(ctx aws.Context, input *PutEventsInput, opts ...request.Option) (*PutEventsOutput, error) {
	req, out := c.PutEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Represents user interaction event information sent using the PutEvents API.
type Event struct {
	_ struct{} `type:"structure"`

	// An ID associated with the event. If an event ID is not provided, Amazon Personalize
	// generates a unique ID for the event. An event ID is not used as an input
	// to the model. Amazon Personalize uses the event ID to distinquish unique
	// events. Any subsequent events after the first with the same event ID are
	// not used in model training.
	EventId *string `locationName:"eventId" min:"1" type:"string"`

	// The type of event. This property corresponds to the EVENT_TYPE field of the
	// Interactions schema.
	//
	// EventType is a required field
	EventType *string `locationName:"eventType" min:"1" type:"string" required:"true"`

	// A string map of event-specific data that you might choose to record. For
	// example, if a user rates a movie on your site, you might send the movie ID
	// and rating, and the number of movie ratings made by the user.
	//
	// Each item in the map consists of a key-value pair. For example,
	//
	// {"itemId": "movie1"}
	//
	// {"itemId": "movie2", "eventValue": "4.5"}
	//
	// {"itemId": "movie3", "eventValue": "3", "numberOfRatings": "12"}
	//
	// The keys use camel case names that match the fields in the Interactions schema.
	// The itemId and eventValue keys correspond to the ITEM_ID and EVENT_VALUE
	// fields. In the above example, the eventType might be 'MovieRating' with eventValue
	// being the rating. The numberOfRatings would match the 'NUMBER_OF_RATINGS'
	// field defined in the Interactions schema.
	//
	// Properties is a required field
	Properties aws.JSONValue `locationName:"properties" type:"jsonvalue" required:"true"`

	// The timestamp on the client side when the event occurred.
	//
	// SentAt is a required field
	SentAt *time.Time `locationName:"sentAt" type:"timestamp" required:"true"`
}

// String returns the string representation
func (s Event) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Event) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Event) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "Event"}
	if s.EventId != nil && len(*s.EventId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventId", 1))
	}
	if s.EventType == nil {
		invalidParams.Add(request.NewErrParamRequired("EventType"))
	}
	if s.EventType != nil && len(*s.EventType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventType", 1))
	}
	if s.Properties == nil {
		invalidParams.Add(request.NewErrParamRequired("Properties"))
	}
	if s.SentAt == nil {
		invalidParams.Add(request.NewErrParamRequired("SentAt"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEventId sets the EventId field's value.
func (s *Event) SetEventId(v string) *Event {
	s.EventId = &v
	return s
}

// SetEventType sets the EventType field's value.
func (s *Event) SetEventType(v string) *Event {
	s.EventType = &v
	return s
}

// SetProperties sets the Properties field's value.
func (s *Event) SetProperties(v aws.JSONValue) *Event {
	s.Properties = v
	return s
}

// SetSentAt sets the SentAt field's value.
func (s *Event) SetSentAt(v time.Time) *Event {
	s.SentAt = &v
	return s
}

// Provide a valid value for the field or parameter.
type InvalidInputException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s InvalidInputException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvalidInputException) GoString() string {
	return s.String()
}

func newErrorInvalidInputException(v protocol.ResponseMetadata) error {
	return &InvalidInputException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InvalidInputException) Code() string {
	return "InvalidInputException"
}

// Message returns the exception's message.
func (s *InvalidInputException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InvalidInputException) OrigErr() error {
	return nil
}

func (s *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InvalidInputException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InvalidInputException) RequestID() string {
	return s.RespMetadata.RequestID
}

type PutEventsInput struct {
	_ struct{} `type:"structure"`

	// A list of event data from the session.
	//
	// EventList is a required field
	EventList []*Event `locationName:"eventList" min:"1" type:"list" required:"true"`

	// The session ID associated with the user's visit.
	//
	// SessionId is a required field
	SessionId *string `locationName:"sessionId" min:"1" type:"string" required:"true"`

	// The tracking ID for the event. The ID is generated by a call to the CreateEventTracker
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateEventTracker.html)
	// API.
	//
	// TrackingId is a required field
	TrackingId *string `locationName:"trackingId" min:"1" type:"string" required:"true"`

	// The user associated with the event.
	UserId *string `locationName:"userId" min:"1" type:"string"`
}

// String returns the string representation
func (s PutEventsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutEventsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PutEventsInput"}
	if s.EventList == nil {
		invalidParams.Add(request.NewErrParamRequired("EventList"))
	}
	if s.EventList != nil && len(s.EventList) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EventList", 1))
	}
	if s.SessionId == nil {
		invalidParams.Add(request.NewErrParamRequired("SessionId"))
	}
	if s.SessionId != nil && len(*s.SessionId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SessionId", 1))
	}
	if s.TrackingId == nil {
		invalidParams.Add(request.NewErrParamRequired("TrackingId"))
	}
	if s.TrackingId != nil && len(*s.TrackingId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("TrackingId", 1))
	}
	if s.UserId != nil && len(*s.UserId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("UserId", 1))
	}
	if s.EventList != nil {
		for i, v := range s.EventList {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EventList", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEventList sets the EventList field's value.
func (s *PutEventsInput) SetEventList(v []*Event) *PutEventsInput {
	s.EventList = v
	return s
}

// SetSessionId sets the SessionId field's value.
func (s *PutEventsInput) SetSessionId(v string) *PutEventsInput {
	s.SessionId = &v
	return s
}

// SetTrackingId sets the TrackingId field's value.
func (s *PutEventsInput) SetTrackingId(v string) *PutEventsInput {
	s.TrackingId = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *PutEventsInput) SetUserId(v string) *PutEventsInput {
	s.UserId = &v
	return s
}

type PutEventsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutEventsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutEventsOutput) GoString() string {
	return s.String()
}

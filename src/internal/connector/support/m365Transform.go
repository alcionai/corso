package support

import (
	"fmt"
	"strconv"
	"strings"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
)

var (
	eventResponsableFields = []string{"responseType"}
	eventRequestableFields = []string{"allowNewTimeProposals", "meetingRequestType", "responseRequested"}
)

// CloneMessageableFields places data from original data into new message object.
// SingleLegacyValueProperty is not populated during this operation
func CloneMessageableFields(orig, message models.Messageable) models.Messageable {
	message.SetAttachments(orig.GetAttachments())
	message.SetAdditionalData(orig.GetAdditionalData())
	message.SetBccRecipients(orig.GetBccRecipients())
	message.SetBody(orig.GetBody())
	message.SetBodyPreview(orig.GetBodyPreview())
	message.SetCcRecipients(orig.GetCcRecipients())
	message.SetConversationId(orig.GetConversationId())
	message.SetConversationIndex(orig.GetConversationIndex())
	message.SetExtensions(orig.GetExtensions())
	message.SetFlag(orig.GetFlag())
	message.SetFrom(orig.GetFrom())
	message.SetHasAttachments(orig.GetHasAttachments())
	message.SetImportance(orig.GetImportance())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetInternetMessageHeaders(orig.GetInternetMessageHeaders())
	message.SetInternetMessageId(orig.GetInternetMessageId())
	message.SetIsDeliveryReceiptRequested(orig.GetIsDeliveryReceiptRequested())
	message.SetIsDraft(orig.GetIsDraft())
	message.SetIsRead(orig.GetIsRead())
	message.SetIsReadReceiptRequested(orig.GetIsReadReceiptRequested())
	message.SetMultiValueExtendedProperties(orig.GetMultiValueExtendedProperties())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetReceivedDateTime(orig.GetReceivedDateTime())
	message.SetReplyTo(orig.GetReplyTo())
	message.SetSender(orig.GetSender())
	message.SetSentDateTime(orig.GetSentDateTime())
	message.SetSubject(orig.GetSubject())
	message.SetToRecipients(orig.GetToRecipients())
	message.SetUniqueBody(orig.GetUniqueBody())
	message.SetWebLink(orig.GetWebLink())

	return message
}

func ToMessage(orig models.Messageable) models.Messageable {
	message := models.NewMessage()
	temp := CloneMessageableFields(orig, message)

	aMessage, ok := temp.(*models.Message)
	if !ok {
		return nil
	}

	return aMessage
}

func SetEventMessageRequest(orig models.Messageable, adtl map[string]any) (models.EventMessageRequestable, error) {
	aMessage := models.NewEventMessageRequest()
	temp := CloneMessageableFields(orig, aMessage)

	message, ok := temp.(models.EventMessageRequestable)
	if !ok {
		return nil, errors.New(*orig.GetId() + " failed to convert to eventMessageRequestable")
	}

	newMessage, err := SetAdditionalDataToEventMessage(adtl, message)
	if err != nil {
		return nil, errors.Wrap(err, *orig.GetId()+" eventMessageRequest could not set additional data")
	}

	additional, err := buildMapFromAdditional(eventRequestableFields, adtl)
	if err != nil {
		return nil, errors.Wrap(err, *orig.GetId()+" eventMessageRequest failed on method buildMapFromAdditional")
	}

	message, ok = newMessage.(models.EventMessageRequestable)
	if !ok {
		return nil, errors.New(*orig.GetId() + " failed to convert to eventMessageRequestable")
	}

	eventMessage, err := setEventRequestableFields(message, additional)
	if err != nil {
		return nil, errors.Wrap(err, "unable to set all fields for eventMessageRequestable")
	}

	return eventMessage, nil
}

func SetEventMessageResponse(orig models.Messageable, adtl map[string]any) (models.EventMessageResponseable, error) {
	aMessage := models.NewEventMessageResponse()
	temp := CloneMessageableFields(orig, aMessage)

	message, ok := temp.(models.EventMessageResponseable)
	if !ok {
		return nil, errors.New(*orig.GetId() + " failed to convert to eventMessageRequestable")
	}

	newMessage, err := SetAdditionalDataToEventMessage(adtl, message)
	if err != nil {
		return nil, errors.Wrap(err, *orig.GetId()+" eventMessageResponse could not set additional data")
	}

	message, ok = newMessage.(models.EventMessageResponseable)
	if !ok {
		return nil, errors.New("unable to create event message responseable from " + *orig.GetId())
	}

	additional, err := buildMapFromAdditional(eventResponsableFields, adtl)
	if err != nil {
		return nil, errors.Wrap(err, *orig.GetId()+" eventMessageResponse failed on method buildMapFromAdditional")
	}

	for key, val := range additional {
		switch key {
		case "responseType":
			temp, err := models.ParseResponseType(*val)
			if err != nil {
				return nil, errors.Wrap(err, *orig.GetId()+"failure to parse response type")
			}

			rType, ok := temp.(*models.ResponseType)
			if !ok {
				return nil, fmt.Errorf(
					"%s : responseType not returned from models.ParseResponseType: %v\t%T",
					*orig.GetId(),
					temp,
					temp,
				)
			}

			message.SetResponseType(rType)

		default:
			return nil, errors.New(key + " not supported for setEventMessageResponse")
		}
	}

	return message, nil
}

// ConvertFromMessageable temporary function. Converts incorrect cast of messageable object to known
// type until upstream can make the appropriate changes
func ConvertFromMessageable(adtl map[string]any, orig models.Messageable) (models.EventMessageable, error) {
	var aType string

	aPointer, ok := adtl["@odata.type"]
	if !ok {
		return nil, errors.New("unknown data type: no @odata.type field")
	}

	ptr, ok := aPointer.(*string)
	if !ok {
		return nil, errors.New("unknown map type encountered")
	}

	aType = *ptr
	if aType == "#microsoft.graph.eventMessageRequest" {
		eventRequest, err := SetEventMessageRequest(orig, adtl)
		if err != nil {
			return nil, err
		}

		eventRequest.SetId(orig.GetId())

		return eventRequest, err
	}

	if aType == "#microsoft.graph.eventMessageResponse" {
		eventMessage, err := SetEventMessageResponse(orig, adtl)
		if err != nil {
			return nil, err
		}

		eventMessage.SetId(orig.GetId())

		return eventMessage, nil
	}

	return nil, errors.New("unknown data type: " + aType)
}

// buildMapFromAdditional returns a submap of map[string]*string from map[string]any
func buildMapFromAdditional(list []string, adtl map[string]any) (map[string]*string, error) {
	returnMap := make(map[string]*string)

	for _, entry := range list {
		ptr, ok := adtl[entry]
		if !ok {
			continue
		}

		value, ok := ptr.(*string)
		if !ok {
			boolConvert, ok := ptr.(*bool)
			if !ok {
				return nil, errors.New("unsupported value type: key: " + entry + fmt.Sprintf(" with type: %T", ptr))
			}

			aBool := *boolConvert
			boolString := strconv.FormatBool(aBool)
			returnMap[entry] = &boolString

			continue
		}

		returnMap[entry] = value
	}

	return returnMap, nil
}

func setEventRequestableFields(
	em models.EventMessageRequestable,
	adtl map[string]*string,
) (models.EventMessageRequestable, error) {
	for key, value := range adtl {
		switch key {
		case "meetingRequestType":
			temp, err := models.ParseMeetingRequestType(*value)
			if err != nil {
				return nil, errors.Wrap(err, *em.GetId()+": failed on models.ParseMeetingRequestType")
			}

			rType, ok := temp.(*models.MeetingRequestType)
			if !ok {
				return nil, errors.New(*em.GetId() + ": failed to set meeting request type")
			}

			em.SetMeetingRequestType(rType)

		case "responseRequested":
			boolValue, err := strconv.ParseBool(*value)
			if err != nil {
				return nil, errors.Wrap(err, *em.GetId()+": failed to set responseRequested")
			}

			em.SetResponseRequested(&boolValue)

		case "allowNewTimeProposals":
			boolValue, err := strconv.ParseBool(*value)
			if err != nil {
				return nil, errors.Wrap(err, *em.GetId()+": failed to set  allowNewTimeProposals")
			}

			em.SetAllowNewTimeProposals(&boolValue)
		}
	}

	return em, nil
}

// SetAdditionalDataToEventMessage sets shared fields for 2 types of EventMessage: Response and Request
func SetAdditionalDataToEventMessage(
	adtl map[string]any,
	newMessage models.EventMessageable,
) (models.EventMessageable, error) {
	for key, entry := range adtl {
		if key == "endDateTime" {
			dateTime := models.NewDateTimeTimeZone()

			mapped, ok := entry.(map[string]*kw.JsonParseNode)
			if ok {
				for key, val := range mapped {
					node := *val

					value, err := node.GetStringValue()
					if err != nil {
						return nil, errors.Wrapf(err, "could not parse string value for %s", key)
					}

					switch key {
					case "dateTime":
						dateTime.SetDateTime(value)
					case "timeZone":
						dateTime.SetTimeZone(value)
					default:
						return nil, errors.New("key not supported DateTime")
					}

					newMessage.SetEndDateTime(dateTime)
				}

				continue
			}
		}

		if key == "startDateTime" {
			dateTime := models.NewDateTimeTimeZone()

			mapped, ok := entry.(map[string]*kw.JsonParseNode)
			if ok {
				for key, val := range mapped {
					node := *val

					value, err := node.GetStringValue()
					if err != nil {
						return nil, errors.Wrapf(err, "could not parse string value for %s", key)
					}

					switch key {
					case "dateTime":
						dateTime.SetDateTime(value)
					case "timeZone":
						dateTime.SetTimeZone(value)
					default:
						return nil, errors.New("key not supported DateTime")
					}

					newMessage.SetStartDateTime(dateTime)
				}

				continue
			}
		}

		if key == "location" {
			aLocation := models.NewLocation()

			mapped, ok := entry.(map[string]*kw.JsonParseNode)
			if ok {
				for key, val := range mapped {
					node := *val

					value, err := node.GetStringValue()
					if err != nil {
						return nil, errors.Wrapf(err, "could not parse string value for key %s", key)
					}

					switch key {
					case "displayName":
						aLocation.SetDisplayName(value)
					case "locationType":
						temp, err := models.ParseLocationType(*value)
						if err != nil {
							return nil, errors.New("location type parse failure")
						}

						lType, ok := temp.(*models.LocationType)
						if !ok {
							return nil, errors.New("location type interface failure")
						}

						aLocation.SetLocationType(lType)
					}
				}
			}

			newMessage.SetLocation(aLocation)
		}

		value, ok := entry.(*string)
		if ok {
			switch key {
			case "isAllDay":
				boolValue, err := strconv.ParseBool(*value)
				if err != nil {
					return nil, err
				}

				newMessage.SetIsAllDay(&boolValue)

			case "isDelegated":
				boolValue, err := strconv.ParseBool(*value)
				if err != nil {
					return nil, err
				}

				newMessage.SetIsDelegated(&boolValue)

			case "isOutOfDate":
				boolValue, err := strconv.ParseBool(*value)
				if err != nil {
					return nil, err
				}

				newMessage.SetIsOutOfDate(&boolValue)

			case "meetingMessageType":
				temp, err := models.ParseMeetingMessageType(*value)
				if err != nil {
					return nil, err
				}

				mType, ok := temp.(*models.MeetingMessageType)
				if !ok {
					return nil, errors.New("failed to create meeting message type")
				}

				newMessage.SetMeetingMessageType(mType)
			}
		}
	}

	return newMessage, nil
}

// ToEventSimplified transforms an event to simplifed restore format
// To overcome some of the MS Graph API challenges, the event object is modified in the following ways:
//   - Instead of adding attendees and generating spurious notifications,
//     add a summary of attendees at the beginning to the event before the original body content
//   - event.attendees is set to an empty list
func ToEventSimplified(orig models.Eventable) models.Eventable {
	attendees := FormatAttendees(orig, *orig.GetBody().GetContentType() == models.HTML_BODYTYPE)
	orig.SetAttendees([]models.Attendeeable{})
	origBody := orig.GetBody()
	newContent := insertStringToBody(origBody, attendees)
	newBody := models.NewItemBody()
	newBody.SetContentType(origBody.GetContentType())
	newBody.SetAdditionalData(origBody.GetAdditionalData())
	newBody.SetOdataType(origBody.GetOdataType())
	newBody.SetContent(&newContent)
	orig.SetBody(newBody)

	return orig
}

type getContenter interface {
	GetContent() *string
	GetContentType() *models.BodyType
}

// insertStringToBody helper function to insert text into models.bodyable
// @returns string containing the content string of altered body.
func insertStringToBody(body getContenter, newContent string) string {
	if body.GetContent() == nil ||
		len(*body.GetContent()) == 0 ||
		body.GetContentType() == nil {
		return ""
	}

	content := *body.GetContent()

	switch *body.GetContentType() {
	case models.TEXT_BODYTYPE:
		return newContent + content

	case models.HTML_BODYTYPE:
		arr := strings.Split(content, "<body>")
		if len(arr) < 2 {
			// malformed html; can't be sure where to insert attendees.
			return newContent + content
		}

		prefix := arr[0] + "<body>"
		interior := arr[1]
		splitOnCloseAngle := strings.Split(interior, ">")

		if len(splitOnCloseAngle) < 3 {
			// no inner elements in body, just insert the new content
			return prefix + newContent + strings.Join(arr[1:], "")
		}

		prefix += splitOnCloseAngle[0] + ">"
		suffix := strings.Join(splitOnCloseAngle[1:], ">")

		return prefix + newContent + suffix
	}

	return newContent + content
}

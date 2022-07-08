package support

import (
	"errors"
	"fmt"
	"strconv"

	kw "github.com/microsoft/kiota-serialization-json-go"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ToMessage transfers all data from old message to new
// message except for the messageId. Required for Restore Operation
func ToMessage(orig models.Messageable) *models.Message {
	message := models.NewMessage()
	message.SetSubject(orig.GetSubject())
	message.SetBodyPreview(orig.GetBodyPreview())
	message.SetBody(orig.GetBody())
	message.SetSentDateTime(orig.GetSentDateTime())
	message.SetReceivedDateTime(orig.GetReceivedDateTime())
	message.SetToRecipients(orig.GetToRecipients())
	message.SetSender(orig.GetSender())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetBccRecipients(orig.GetBccRecipients())
	message.SetCcRecipients(orig.GetCcRecipients())
	message.SetReplyTo(orig.GetReplyTo())
	message.SetFlag(orig.GetFlag())
	message.SetHasAttachments(orig.GetHasAttachments())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetConversationId(orig.GetConversationId())
	message.SetExtensions(orig.GetExtensions())
	message.SetFlag(orig.GetFlag())
	message.SetFrom(orig.GetFrom())
	message.SetImportance(orig.GetImportance())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetInternetMessageId(orig.GetInternetMessageId())
	message.SetInternetMessageHeaders(orig.GetInternetMessageHeaders())
	message.SetIsDeliveryReceiptRequested(orig.GetIsDeliveryReceiptRequested())
	message.SetIsRead(orig.GetIsRead())
	message.SetIsReadReceiptRequested(orig.GetIsReadReceiptRequested())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetMultiValueExtendedProperties(orig.GetMultiValueExtendedProperties())
	message.SetUniqueBody(orig.GetUniqueBody())
	message.SetWebLink(orig.GetWebLink())
	return message

}

func SetEventMessageRequest(orig models.Messageable) *models.EventMessageRequest {
	message := models.NewEventMessageRequest()
	message.SetSubject(orig.GetSubject())
	message.SetBodyPreview(orig.GetBodyPreview())
	message.SetBody(orig.GetBody())
	message.SetSentDateTime(orig.GetSentDateTime())
	message.SetReceivedDateTime(orig.GetReceivedDateTime())
	message.SetToRecipients(orig.GetToRecipients())
	message.SetSender(orig.GetSender())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetBccRecipients(orig.GetBccRecipients())
	message.SetCcRecipients(orig.GetCcRecipients())
	message.SetReplyTo(orig.GetReplyTo())
	message.SetFlag(orig.GetFlag())
	message.SetHasAttachments(orig.GetHasAttachments())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetConversationId(orig.GetConversationId())
	message.SetExtensions(orig.GetExtensions())
	message.SetFlag(orig.GetFlag())
	message.SetFrom(orig.GetFrom())
	message.SetImportance(orig.GetImportance())
	message.SetInferenceClassification(orig.GetInferenceClassification())
	message.SetInternetMessageId(orig.GetInternetMessageId())
	message.SetInternetMessageHeaders(orig.GetInternetMessageHeaders())
	message.SetIsDeliveryReceiptRequested(orig.GetIsDeliveryReceiptRequested())
	message.SetIsRead(orig.GetIsRead())
	message.SetIsReadReceiptRequested(orig.GetIsReadReceiptRequested())
	message.SetParentFolderId(orig.GetParentFolderId())
	message.SetMultiValueExtendedProperties(orig.GetMultiValueExtendedProperties())
	message.SetUniqueBody(orig.GetUniqueBody())
	message.SetWebLink(orig.GetWebLink())
	return message
}

// ConvertFromMessageable temporary function. Converts incorrect cast of messageable object to known
// type until upstream can make the appropriate changes
func ConvertFromMessageable(adtl map[string]any, orig models.Messageable) (models.Messageable, error) {
	// Verify if it is a known type
	aPointer, ok := adtl["@odata.type"]
	if ok {
		ptr, ok := aPointer.(*string)
		if !ok {
			return nil, errors.New("unknown map type encountered")
		}
		if *ptr == "#microsoft.graph.eventMessageRequest" {
			newMessage := SetEventMessageRequest(orig)
			newMessage.SetId(orig.GetId())
			for key, entry := range adtl {
				if key == "endDateTime" {
					dateTime := models.NewDateTimeTimeZone()
					mapped, ok := entry.(map[string]*kw.JsonParseNode)
					if ok {
						for key, val := range mapped {
							node := *val
							value, err := node.GetStringValue()
							if err != nil {
								return nil, err
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
								return nil, err
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
					fmt.Printf("%s of type %T\n", key, entry)
					aLocation := models.NewLocation()
					mapped, ok := entry.(map[string]*kw.JsonParseNode)
					if ok {
						for key, val := range mapped {
							node := *val
							value, err := node.GetStringValue()
							if err != nil {
								fmt.Printf("Err: %v\n", err)
							}
							switch key {
							case "displayName":
								aLocation.SetDisplayName(value)
							case "locationType":
								ty, err := models.ParseLocationType(*value)
								fmt.Printf("What %T\n", ty)
								if err != nil {
									return nil, errors.New("location type parse failure")
								}
								lType, ok := ty.(*models.LocationType)
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
					case "meetingRequestType":
						temp, err := models.ParseMeetingRequestType(*value)
						if err != nil {
							return nil, err
						}
						rType, ok := temp.(*models.MeetingRequestType)
						if !ok {
							return nil, errors.New("failed to create request type")
						}
						newMessage.SetMeetingRequestType(rType)
					}

				}
			}
			return newMessage, nil
			//Time to set additional data
		}
	}
	return nil, errors.New("unknown data type")
}

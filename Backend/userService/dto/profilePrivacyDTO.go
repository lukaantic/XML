package dto

import "userService/model"

type ProfilePrivacyDTO struct {
	Id                  string            `json:"_id"`
	PrivacyType         model.PrivacyType `json:"privacyType"`
	AllMessagesRequests bool              `json:"allMessageRequests"`
}

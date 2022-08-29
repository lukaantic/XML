package dto

import "postService/model"

type NotificationDTO struct {
	PostId           string                 `json:"PostId"`
	Text             string                 `json:"text"`
	UserId           string                 `json:"userId"`
	NotificationType model.NotificationType `json:"notificationType"`
}

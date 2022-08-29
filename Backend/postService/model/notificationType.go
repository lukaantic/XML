package model

type NotificationType int

const (
	CommentNotification NotificationType = iota
	LikeNotification
	DislikeNotification
)

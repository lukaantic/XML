package model

type ProfilePrivacy struct {
	PrivacyType        PrivacyType `bson:"privacyType, omitempty"`
	AllMessageRequests bool        `bson:"allMessageRequests, omitempty"`
}

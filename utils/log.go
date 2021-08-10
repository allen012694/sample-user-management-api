package utils

import "fmt"

type ActivityLog struct {
	Subject   string `json:"subject"`
	SubjectId string `json:"subject_id"`

	Object   string `json:"object"`
	ObjectId string `json:"object_id"`

	Action string `json:"action"`

	// json string
	Request  string `json:"request"`
	Response string `json:"response"`
}

func (l ActivityLog) String() string {
	return fmt.Sprintf("%v (%v) %v %v (%v) | req: %v | res: %v",
		l.Subject, l.SubjectId,
		l.Action,
		l.Object, l.ObjectId,
		l.Request, l.Response)
}

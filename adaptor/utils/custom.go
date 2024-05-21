package utils

type Recipient struct {
	ChatId string
}

func (r *Recipient) Recipient() string {
	return r.ChatId
}

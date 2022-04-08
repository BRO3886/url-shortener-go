package shortener

import "strconv"

type Redirect struct {
	Code      string `json:"code" bson:"code"`
	URL       string `json:"url" bson:"url" validate:"required,url"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}

func (r *Redirect) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":       r.Code,
		"url":        r.URL,
		"created_at": r.CreatedAt,
	}
}

func (r *Redirect) FromMap(data map[string]string) error {
	created, err := strconv.ParseInt(data["created_at"], 10, 64)
	if err != nil {
		return err
	}

	r.Code = data["code"]
	r.URL = data["url"]
	r.CreatedAt = created

	return nil
}

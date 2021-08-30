package runtime

// Token is token transform between element.
type Token interface {
	ID() string
	SetMetadata(map[string]interface{}) error
}

type Form struct {
	FID      string                 `json:"id"`
	Metadata map[string]interface{} `json:"metadata"`
}

func NewForm() *Form {
	return &Form{}
}

func (f *Form) ID() string {
	return f.FID
}

func (f *Form) SetMetadata(metedata map[string]interface{}) error {
	return nil
}

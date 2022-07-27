package validations

type GetBlogsValidation struct {
	Title  string `json:"name" validate:"required,min=1,max=16"`
	Author string `json:"author" validate:"required,min=1,max=16"`
	Url    string `json:"url" validate:"required,min=1,max=256"`
	Likes  int    `json:"likes" validate:"min=0"`
}

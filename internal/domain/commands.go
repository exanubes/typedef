package domain

type GenerateCommandInput struct {
	Input      string
	InputType  string
	Format     string
	Target     string
	OutputPath string
}

type GenerateCommandOutput struct {
	Code   string `json:"code"`
	Format Format `json:"format"`
}

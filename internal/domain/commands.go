package domain

type GenerateCommandInput struct {
	Input      string `json:"input"`
	InputType  string `json:"input_type"`
	Format     Format `json:"format"`
	Target     string `json:"target"`
	OutputPath string `json:"output_path"`
}

type GenerateCommandOutput struct {
	Code   string `json:"code"`
	Format Format `json:"format"`
}

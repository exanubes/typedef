package domain

type GenerateCommand struct {
	Input      string
	Format     string
	Target     string
	OutputPath string
}

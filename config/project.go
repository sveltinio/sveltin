package config

// ProjectConfig represents the user selections when creating a new sveltin project.
type ProjectConfig struct {
	ProjectName   string
	CSSLibName    string
	ThemeName     string
	NPMClientName string
}

// NewProjectConfig  creates a pointer to a ProjectConfig struct.
func NewProjectConfig(name, css, theme, npmc string) *ProjectConfig {
	return &ProjectConfig{
		ProjectName:   name,
		CSSLibName:    css,
		ThemeName:     theme,
		NPMClientName: npmc,
	}
}

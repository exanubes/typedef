package targets

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileTarget struct {
	path string
}

func NewFileTarget(path string) *FileTarget {
	return &FileTarget{
		path: path,
	}
}

func (target *FileTarget) Send(input string) error {
	is_absolute := filepath.IsAbs(target.path)
	var path string
	var err error
	if is_absolute {
		path, err = target.handle_absolute()
	} else {
		path, err = target.handle_relative()
	}

	if err != nil {
		return fmt.Errorf("Failed to resolve path for %s:\n| %w", target.path, err)
	}

	err = os.MkdirAll(filepath.Dir(path), 0o755)

	if err != nil {
		return fmt.Errorf("Failed to create directories for %s:\n| %w", path, err)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)

	if err != nil {
		return fmt.Errorf("Failed to open file at location %s:\n| %w", path, err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(input)); err != nil {
		return fmt.Errorf("Failed to write data to file:\n| %w", err)
	}

	fmt.Println("\nOutput saved to: ", path)

	return nil
}

func (target *FileTarget) handle_absolute() (string, error) {
	return target.path, nil
}
func (target *FileTarget) handle_relative() (string, error) {
	dir, err := os.Getwd()

	if err != nil {
		return "", err
	}
	return filepath.Join(dir, target.path), nil
}

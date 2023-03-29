package helper

import (
	"os"

	"github.com/makves-test-task/config"
	"github.com/makves-test-task/pkg/logger"
)

type FileManager struct {
	File *os.File
	Log  *logger.Logger
	Cfg  *config.Config
}

func New(log *logger.Logger, cfg *config.Config) (*FileManager, error) {
	f, err := os.Open(cfg.FilePath)
	if err != nil {
		return &FileManager{}, err
	}

	return &FileManager{
		File: f,
		Log:  log,
		Cfg:  cfg,
	}, nil
}

func (f *FileManager) ReOpenFile() error {
	file, err := os.Open(f.Cfg.FilePath)
	if err != nil {
		return err
	}
	f.File = file
	return nil
}

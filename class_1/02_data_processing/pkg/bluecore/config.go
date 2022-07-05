package bluecore

import "github.com/spf13/afero"

type Config struct {
	FilePath string
	Fs       afero.Fs
}

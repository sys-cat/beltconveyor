package beltconveyor

import (
	"image"
	"os"
)

type (
	Beltconveyor struct {
		Config    image.Config
		Extension string
		ReadPath  string
		File      *os.File
	}
)

func New(path string) Beltconveyor {
	return Beltconveyor{
		ReadPath: path,
	}
}

func (b *Beltconveyor) Decode() (err error) {
	b.File, err = os.Open(b.ReadPath)
	if err != nil {
		return err
	}
	defer b.File.Close()
	b.Config, b.Extension, err = image.DecodeConfig(b.File)
	return err
}

func (b *Beltconveyor) ConvertPNG() (err error) {
	return err
}

func (b *Beltconveyor) ConvertJPG() (err error) {
	return err
}

func (b *Beltconveyor) ConvertWebp() (err error) {
	return err
}

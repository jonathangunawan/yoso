package yoso

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const (
	errIneffective = "Ineffective Dependency Parameter"
	logPrefix      = "Yoso"
)

func NewWriter(cfg Config) (*Dep, error) {
	if (cfg.UsePart && cfg.LimitPerPart <= 0) || (!cfg.UsePart && cfg.LimitPerPart > 0) {
		return nil, errors.Wrap(fmt.Errorf(errIneffective), logPrefix)
	}

	f, w, resultFile, err := newFile(cfg, 1)
	if err != nil {
		return nil, errors.Wrap(err, logPrefix)
	}

	return &Dep{
		File:        f,
		Writer:      w,
		Cfg:         cfg,
		PartCounter: 1,
		ResultFiles: []string{resultFile},
	}, nil
}

func newFile(cfg Config, partCounter int) (*os.File, *csv.Writer, string, error) {
	completeFileName := fileNaming(cfg, partCounter)
	f, err := os.Create(completeFileName)
	if err != nil {
		return nil, nil, "", err
	}

	//init csv writer dependency
	w := csv.NewWriter(f)

	//change separator
	w.Comma = cfg.Separator

	//write header
	if len(cfg.Header) > 0 {
		err := w.Write(cfg.Header)
		if err != nil {
			return nil, nil, "", err
		}
	}

	return f, w, completeFileName, nil
}

func fileNaming(cfg Config, partCounter int) string {
	namePath := cfg.Path + cfg.FileName

	if cfg.UsePart {
		namePath += fmt.Sprintf("_part%d", partCounter)
	}

	return namePath + ".csv"
}

func (d *Dep) Write(input []string) error {
	if d.Cfg.LimitPerPart > 0 && d.Cfg.LimitPerPart == d.LimitCounter {
		//reset limitCounter
		d.LimitCounter = 0

		//add partCounter
		d.PartCounter++

		//flush and close file
		err := d.Close()
		if err != nil {
			return errors.Wrap(err, logPrefix)
		}

		f, w, resultFile, err := newFile(d.Cfg, d.PartCounter)
		if err != nil {
			return errors.Wrap(err, logPrefix)
		}

		d.File = f
		d.Writer = w
		d.ResultFiles = append(d.ResultFiles, resultFile)
	}

	err := d.Writer.Write(input)
	if err != nil {
		return errors.Wrap(err, logPrefix)
	}

	//just count even UsePart is true
	d.LimitCounter++

	return nil
}

func (d *Dep) GetResultFiles() []string {
	return d.ResultFiles
}

func (d *Dep) Close() error {
	d.Writer.Flush()
	return d.File.Close()
}

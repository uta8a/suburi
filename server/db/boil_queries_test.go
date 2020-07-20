// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"regexp"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

var dbNameRand *rand.Rand

func MustTx(transactor boil.ContextTransactor, err error) boil.ContextTransactor {
	if err != nil {
		panic(fmt.Sprintf("Cannot create a transactor: %s", err))
	}
	return transactor
}

func newFKeyDestroyer(regex *regexp.Regexp, reader io.Reader) io.Reader {
	return &fKeyDestroyer{
		reader: reader,
		rgx:    regex,
	}
}

type fKeyDestroyer struct {
	reader io.Reader
	buf    *bytes.Buffer
	rgx    *regexp.Regexp
}

func (f *fKeyDestroyer) Read(b []byte) (int, error) {
	if f.buf == nil {
		all, err := ioutil.ReadAll(f.reader)
		if err != nil {
			return 0, err
		}

		all = bytes.Replace(all, []byte{'\r', '\n'}, []byte{'\n'}, -1)
		all = f.rgx.ReplaceAll(all, []byte{})
		f.buf = bytes.NewBuffer(all)
	}

	return f.buf.Read(b)
}

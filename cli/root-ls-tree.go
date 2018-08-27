/*
 * Copyright (C) 2018  CERN for the benefit of the LHCb collaboration
 * Author: Paul Seyfert <pseyfert@cern.ch>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * In applying this licence, CERN does not waive the privileges and immunities
 * granted to it by virtue of its status as an Intergovernmental Organization
 * or submit itself to any jurisdiction.
 */

// heavily inspired by go-hep's root-ls and root_numpy's list_trees

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"go-hep.org/x/hep/rootio"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`Usage: root-ls-tree file1.root

prints full paths of contained trees.`)
	}

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "error: you need to give a ROOT file\n\n")
		flag.Usage()
		os.Exit(1)
	}

	fname := flag.Args()[0]

	f, err := rootio.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "rootio: failed to open [%s]: %v\n", fname, err)
		os.Exit(1)
	}
	defer f.Close()

	var iowriter bytes.Buffer
	w := fullpathWriter{pth: []byte(""), w: &iowriter}
	for _, k := range f.Keys() {
		walk(w, k)
	}
	fmt.Printf("%s\n", iowriter.String())
}

func walk(pth fullpathWriter, k rootio.Key) {
	obj := k.Value()
	_, ok := obj.(rootio.Tree)
	if ok {
		fmt.Fprintf(&pth, "%s\t", k.Name())
		return
	}
	if dir, ok := obj.(rootio.Directory); ok {
		w := newSubdir([]byte(k.Name()+"/"), pth)
		for _, k := range dir.Keys() {
			walk(*w, k)
		}
	}
}

type fullpathWriter struct {
	pth []byte
	w   io.Writer
}

func newSubdir(dir []byte, w fullpathWriter) *fullpathWriter {
	return &fullpathWriter{
		pth: dir,
		w:   &w,
	}
}

func (w *fullpathWriter) Write(data []byte) (int, error) {
	return w.w.Write(append(w.pth, data...))
}

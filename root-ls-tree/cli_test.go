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

// testing the cli output

package main

import (
	"bytes"
	"testing"

	"go-hep.org/x/hep/rootio"
)

var benchmark_result string

func wrap_walk_for_test(w fullpathWriter, tb testing.TB, fname string) {
	f, err := rootio.Open(fname)
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	walk_file(w, f)
}

func TestBigFileCLI(t *testing.T) {
	var iowriter bytes.Buffer
	w := fullpathWriter{pth: []byte(""), w: &iowriter}

	fname := "../testdata/TMVA.root"
	wrap_walk_for_test(w, t, fname)

	reference := []byte("dataset/Method_BDTG/BDTG/MonitorNtuple	dataset/Method_BDT/BDT/MonitorNtuple	dataset/TestTree	dataset/TrainTree	")
	if 0 != bytes.Compare(iowriter.Bytes(), reference) {
		t.Fatalf("unexpected output.\nExpected:\n%s\nGot:\n%s\n", reference, iowriter.String())
	}
}

func TestSmallFileCLI(t *testing.T) {
	var iowriter bytes.Buffer
	w := fullpathWriter{pth: []byte(""), w: &iowriter}

	fname := "../testdata/sample-6.14.00-uncompressed.root"
	wrap_walk_for_test(w, t, fname)

	reference := []byte("sample	")
	if 0 != bytes.Compare(iowriter.Bytes(), reference) {
		t.Fatalf("unexpected output.\nExpected:\n%s\nGot:\n%s\n", reference, iowriter.String())
	}
}

func BenchmarkSmallCLI(b *testing.B) {
	fname := "../testdata/sample-6.14.00-uncompressed.root"
	for n := 0; n < b.N; n++ {
		var iowriter bytes.Buffer
		w := fullpathWriter{pth: []byte(""), w: &iowriter}

		wrap_walk_for_test(w, b, fname)

		benchmark_result = iowriter.String()
	}
}

func BenchmarkBigCLI(b *testing.B) {
	fname := "../testdata/TMVA.root"
	for n := 0; n < b.N; n++ {
		var iowriter bytes.Buffer
		w := fullpathWriter{pth: []byte(""), w: &iowriter}

		wrap_walk_for_test(w, b, fname)

		benchmark_result = iowriter.String()
	}
}

// Copyright 2017 The go-eubchain Authors
// This file is part of the go-eubchain library.
//
// The go-eubchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-eubchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-eubchain library. If not, see <http://www.gnu.org/licenses/>.

// +build go1.8

package ethash

import "testing"

// Tests whether the dataset size calculator works correctly by cross checking the
// hard coded lookup table with the value generated by it.
func TestSizeCalculations(t *testing.T) {
	// Verify all the cache and dataset sizes from the lookup table.
	for epoch, want := range cacheSizes {
		if size := calcCacheSize(epoch); size != want {
			t.Errorf("cache %d: cache size mismatch: have %d, want %d", epoch, size, want)
		}
	}
	for epoch, want := range datasetSizes {
		if size := calcDatasetSize(epoch); size != want {
			t.Errorf("dataset %d: dataset size mismatch: have %d, want %d", epoch, size, want)
		}
	}
}

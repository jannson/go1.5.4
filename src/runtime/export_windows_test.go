// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Export guts for testing.

package runtime

var TestingWER = &testingWER

func LoadLibraryExStatus() (useEx, haveEx, haveFlags bool) {
	return useLoadLibraryEx, _LoadLibraryExW != nil, _AddDllDirectory != nil
}

// Copyright 2017 The go-shaft Authors
// This file is part of the go-shaft library.
//
// The go-shaft library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-shaft library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-shaft library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

const mainnetAllocData = "\xe2\u151c\x0ea3\xfev\xc7\xc9\x03\xe2\xbe,\xf9\xf7\xa3Y\u007f\xb8\xd7v\x8b\x04\xf6\x8c\xa6\xd8\u0351\xc6\x00\x00\x00"
const testnetAllocData = "\xe2\u151c\x0ea3\xfev\xc7\xc9\x03\xe2\xbe,\xf9\xf7\xa3Y\u007f\xb8\xd7v\x8b\x04\xf6\x8c\xa6\xd8\u0351\xc6\x00\x00\x00"
const rinkebyAllocData = "\xe2\u151c\x0ea3\xfev\xc7\xc9\x03\xe2\xbe,\xf9\xf7\xa3Y\u007f\xb8\xd7v\x8b\x04\xf6\x8c\xa6\xd8\u0351\xc6\x00\x00\x00"
const devAllocData = "\xe2\u151c\x0ea3\xfev\xc7\xc9\x03\xe2\xbe,\xf9\xf7\xa3Y\u007f\xb8\xd7v\x8b\x04\xf6\x8c\xa6\xd8\u0351\xc6\x00\x00\x00"

// +build amd64
// +build !appengine

/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package v1

func haveSSE42() bool
func scanStringSSE(s []byte, j int) (int, byte)

var sse42 = haveSSE42()

func scanString(s []byte, j int) (int, byte) {
	// XXX The following fails to compile on Go 1.2.
	/*
		if false && sse42 {
			return scanStringSSE(s, j)
		}
	*/

	for {
		if j >= len(s) {
			return j, 0
		}

		c := s[j]
		j++
		if byteLookupTable[c]&sliceStringMask == 0 {
			continue
		}

		return j, c
	}
}

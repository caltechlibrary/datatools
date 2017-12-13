//
// datatools package is a collection of Go based command
// line tools for working with JSON content
//
// @Author R. S. Doiel, <rsdoiel@caltech.edu>
//
// Copyright (c) 2017, Caltech
// All rights not granted herein are expressly reserved by Caltech.
//
// Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package datatools

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseRange takes a string in the form of a "range expression" like 1,2 (one and two), 1-3 (one, two, three)
// or 1,2,8-10 (one, two, eight, nine, ten) and returns an array of ints holding the values of the range expression.
func ParseRange(s string) ([]int, error) {
	r := []int{}
	// Convert colon delimited ranges to dashes
	if strings.Contains(s, ":") {
		s = strings.Replace(s, ":", "-", -1)
	}
	cells := strings.Split(s, ",")
	for _, c := range cells {
		if strings.Contains(c, "-") {
			p := strings.Split(c, "-")
			if len(p) != 2 {
				return r, fmt.Errorf("%q is not an int range like 10 - 13", c)
			}
			p[0], p[1] = strings.TrimSpace(p[0]), strings.TrimSpace(p[1])
			start, err := strconv.Atoi(p[0])
			if err != nil {
				return r, fmt.Errorf("%q, %s", p[0], err)
			}
			end, err := strconv.Atoi(p[1])
			if err != nil {
				return r, fmt.Errorf("%q, %s", p[1], err)
			}
			if start > end {
				start, end = end, start
			}
			for i := start; i <= end; i++ {
				r = append(r, i)
			}
		} else {
			i, err := strconv.Atoi(strings.TrimSpace(c))
			if err != nil {
				return r, fmt.Errorf("%q, %s", c, err)
			}
			r = append(r, i)
		}
	}
	return r, nil
}

// Copyright © 2019 The Homeport Team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bunt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/gonvenience/bunt"
)

var _ = Describe("working with colored strings", func() {
	Context("cut a substring from a colored string", func() {
		BeforeEach(func() {
			ColorSetting = ON
		})

		AfterEach(func() {
			ColorSetting = AUTO
		})

		It("should be possible to cut out a substring from a colored string", func() {
			input := "Example: \x1b[1mbold\x1b[0m, \x1b[3mitalic\x1b[0m, \x1b[4munderline\x1b[0m, \x1b[38;2;133;247;7mforeground\x1b[0m, and \x1b[48;2;133;247;7mbackground\x1b[0m."
			result, err := ParseString(input)
			Expect(err).ToNot(HaveOccurred())
			Expect(result).ToNot(BeNil())

			result.Substring(11, 28)

			expected := "\x1b[1mld\x1b[0m, \x1b[3mitalic\x1b[0m, \x1b[4munder\x1b[0m"
			actual := result.String()
			Expect(actual).To(BeEquivalentTo(expected))
		})
	})
})

// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nameLen = 6
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomName generates random names
func RandomName(prefix string) string {
	return strings.ToLower(fmt.Sprintf("%s-%s", prefix, stringWithCharset(nameLen, charset)))
}

// ExtractID extracts ID from location of response
func ExtractID(location string) (int64, error) {
	idstr := location[strings.LastIndex(location, "/")+1:]
	return strconv.ParseInt(idstr, 10, 64)
}

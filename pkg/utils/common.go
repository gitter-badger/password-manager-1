// Copyright © 2019 Thilina Manamgoda
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this fileio except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


package utils

// AESEncryptID is the unique identifier for this encryptor
const AESEncryptID = "AES"

// IsValidByteSlice method check whether the Slice is valid or not
func IsValidByteSlice(data []byte) bool {
	return (data != nil) && (len(data) != 0)
}

// StringSliceContains check whether the specified key is in the String slice
func StringSliceContains(key string, s []string) bool {
	for _, v := range s {
		if key == v {
			return true
		}
	}
	return false
}


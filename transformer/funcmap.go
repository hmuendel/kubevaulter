/*
 * Copyright 2018 Hans Mündelein
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package transformer

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

type FuncMap map[string]func(string) string

func Identity(in string) string  {
	return in
}

func Sha1(in string) string  {
	bytes := sha1.Sum([]byte(in))
	return hex.EncodeToString(bytes[:])
}

func Sha256(in string) string {
	bytes := sha256.Sum256([]byte(in))
	return hex.EncodeToString(bytes[:])
}

func Md5(in string) string {
	bytes := md5.Sum([]byte(in))
	return hex.EncodeToString(bytes[:])
}

func DefaultFuncMap() FuncMap  {
	m := FuncMap{}
	m[""] = Identity
	m["NONE"] = Identity
	m["SHA1"] = Sha1
	m["SHA256"] = Sha256
	m["MD5"] = Md5
	return  m
}

/*
 * Copyright 2020 Anton Globa
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package gotspl

import (
	"bytes"
	"errors"
	"strings"
)

const RUN_NAME = "RUN"

type RunImpl struct {
	file *string
}

type RunBuilder interface {
	TSPLCommand
	File(name string) RunBuilder
}

func RunCmd() RunBuilder {
	return RunImpl{}
}

func (r RunImpl) GetMessage() ([]byte, error) {
	if r.file == nil || len(*r.file) == 0 {
		return nil, errors.New("ParseError RUN Command: file should be specified")
	}

	buf := bytes.NewBufferString(RUN_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString(DOUBLE_QUOTE)
	buf.WriteString(strings.ToUpper(*r.file))
	buf.WriteString(DOUBLE_QUOTE)
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}

func (r RunImpl) File(name string) RunBuilder {
	if r.file == nil {
		r.file = new(string)
	}
	*r.file = name
	return r
}

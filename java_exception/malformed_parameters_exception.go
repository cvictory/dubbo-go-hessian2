/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package java_exception

// MalformedParametersException represents an exception of the same name in java
type MalformedParametersException struct {
	SerialVersionUID     int64
	DetailMessage        string
	SuppressedExceptions []Throwabler
	StackTrace           []StackTraceElement
	Cause                Throwabler
}

// Error output error message
func (e MalformedParametersException) Error() string {
	return e.DetailMessage
}

// JavaClassName  java fully qualified path
func (MalformedParametersException) JavaClassName() string {
	return "java.lang.reflect.MalformedParametersException"
}

// NewMalformedParametersException is the constructor
func NewMalformedParametersException(detailMessage string) *MalformedParametersException {
	return &MalformedParametersException{DetailMessage: detailMessage}
}

// equals to getStackTrace in java
func (e MalformedParametersException) GetStackTrace() []StackTraceElement {
	return e.StackTrace
}
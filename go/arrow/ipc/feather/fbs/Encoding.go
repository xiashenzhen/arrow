// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fbs

type Encoding = int8
const (
	EncodingPLAIN Encoding = 0
	/// Data is stored dictionary-encoded
	/// dictionary size: <INT32 Dictionary size>
	/// dictionary data: <TYPE primitive array>
	/// dictionary index: <INT32 primitive array>
	///
	/// TODO: do we care about storing the index values in a smaller typeclass
	EncodingDICTIONARY Encoding = 1
)

var EnumNamesEncoding = map[Encoding]string{
	EncodingPLAIN:"PLAIN",
	EncodingDICTIONARY:"DICTIONARY",
}


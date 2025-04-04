// Copyright 2019-2025 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package port

import (
	"fmt"
)

// ParsePortRange parses the port range and returns the start and end of the range.
func ParsePortRange(value string) (start, end uint16, err error) {
	_, err = fmt.Sscanf(value, "%d-%d", start, end)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid port range %s", value)
	}

	return start, end, nil
}

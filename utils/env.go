
/**
Copyright © 2020 Appvia Ltd <info@appvia.io>
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

package utils

import (
  "os"
  "fmt"
)

func GetEnv(str string) string {
  envVar := os.Getenv(str)
  return envVar
}

func DefaultToEnv(keyFromEnv string, keyFromFlag string) string {
  if keyFromFlag == "" && keyFromEnv == "" {
    fmt.Println("Set API_KEY as environment variable or specify --api-key flag or -k flag.")
    os.Exit(1)
  } else if keyFromFlag == "" && keyFromEnv != "" {
      return keyFromEnv
  }
  return keyFromFlag
}

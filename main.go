/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

// I followd this tuto : https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177

import "my_calc/cmd"
import "fmt"

func main() {
  fmt.Println("Inside main function")
  cmd.Execute()
}

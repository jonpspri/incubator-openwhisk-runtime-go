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
package openwhisk

import (
	"fmt"
	"io/ioutil"
)

func ExampleNewExecutor_failed() {
	log, _ := ioutil.TempFile("", "log")
	proc := NewExecutor(log, log, "true")
	err := proc.Start()
	fmt.Println(err)
	proc.Stop()
	proc = NewExecutor(log, log, "/bin/pwd")
	err = proc.Start()
	fmt.Println(err)
	proc.Stop()
	proc = NewExecutor(log, log, "donotexist")
	err = proc.Start()
	fmt.Println(err)
	proc.Stop()
	proc = NewExecutor(log, log, "/etc/passwd")
	err = proc.Start()
	fmt.Println(err)
	proc.Stop()
	// Output:
	// command exited
	// command exited
	// command exited
	// command exited
}

func ExampleNewExecutor_bc() {
	log, _ := ioutil.TempFile("", "log")
	proc := NewExecutor(log, log, "_test/bc.sh")
	err := proc.Start()
	fmt.Println(err)
	res, _ := proc.Interact([]byte("2+2"))
	fmt.Printf("%s", res)
	proc.Stop()
	dump(log)
	// Output:
	// <nil>
	// 4
	// XXX_THE_END_OF_A_WHISK_ACTIVATION_XXX
	// XXX_THE_END_OF_A_WHISK_ACTIVATION_XXX
}

func ExampleNewExecutor_hello() {
	log, _ := ioutil.TempFile("", "log")
	proc := NewExecutor(log, log, "_test/hello.sh")
	err := proc.Start()
	fmt.Println(err)
	res, _ := proc.Interact([]byte(`{"value":{"name":"Mike"}}`))
	fmt.Printf("%s", res)
	proc.Stop()
	dump(log)
	// Output:
	// <nil>
	// {"hello": "Mike"}
	// msg=hello Mike
	// XXX_THE_END_OF_A_WHISK_ACTIVATION_XXX
	// XXX_THE_END_OF_A_WHISK_ACTIVATION_XXX
}

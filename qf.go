// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package main

import (
	"fmt"
	"os"

	"github.com/quickfixgo/examples/cmd"
)

func main() {
	//执行了examples中的cmd
	fmt.Println("cmd execute")
	err := cmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

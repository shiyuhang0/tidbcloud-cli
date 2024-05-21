// Copyright 2024 PingCAP, Inc.
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

package iostream

import (
	"bytes"
	"github.com/mattn/go-isatty"
	"io"
	"os"
)

type IOStreams struct {
	Out       io.Writer
	Err       io.Writer
	CanPrompt bool
}

func System() *IOStreams {
	canPrompt := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())

	return &IOStreams{
		Out:       os.Stdout,
		Err:       os.Stderr,
		CanPrompt: canPrompt,
	}
}

func Test() *IOStreams {
	return &IOStreams{
		Out:       &bytes.Buffer{},
		Err:       &bytes.Buffer{},
		CanPrompt: false,
	}
}

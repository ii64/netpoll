// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux

package netpoll

import "github.com/cloudwego/netpoll/uring"

// TODO: init uringPoll
func openIOURingPoll(entries uint32, opts ...uring.SetupOpt) *uringPoll {
	poll := new(uringPoll)
	ring, err := uring.IOURing(entries, opts...)
	if err != nil {
		panic(err)
	}
	poll.fd = ring.Fd()
	return poll
}

// TODO: build uringPoll
type uringPoll struct {
	fd int
}

// TODO: Wait implements Poll.
func (p *uringPoll) Wait() error

// TODO: Close implements Poll.
func (p *uringPoll) Close() error

// TODO: Trigger implements Poll.
func (p *uringPoll) Trigger() error

// TODO: Control implements Poll.
func (p *uringPoll) Control(operator *FDOperator, event PollEvent) error

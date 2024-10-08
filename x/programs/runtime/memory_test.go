// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runtime

import (
	"context"
	_ "embed"
	"os"
	"testing"

	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/bytecodealliance/wasmtime-go/v14"
	"github.com/stretchr/testify/require"
)

var log = logging.NewLogger(
	"",
	logging.NewWrappedCore(
		logging.Info,
		os.Stderr,
		logging.Plain.ConsoleEncoder(),
	))

func TestLimitMaxMemory(t *testing.T) {
	require := require.New(t)

	// memory has a single page
	wasm, err := wasmtime.Wat2Wasm(`
	(module

	  (memory 2) ;; 2 pages
	  (export "memory" (memory 0))
	)
	`)
	require.NoError(err)

	// wasm defines 2 pages of memory but runtime set max 1 page
	maxUnits := uint64(1)
	cfg, err := NewConfigBuilder().
		WithLimitMaxMemory(1 * MemoryPageSize). // 1 page
		Build()
	require.NoError(err)
	runtime := New(log, cfg, nil)
	err = runtime.Initialize(context.Background(), wasm, maxUnits)
	require.ErrorContains(err, "memory minimum size of 2 pages exceeds memory limits")
}

func TestLimitMaxMemoryGrow(t *testing.T) {
	require := require.New(t)

	wasm, err := wasmtime.Wat2Wasm(`
	(module
	
	  (memory 1) ;; 1 pages
	  (export "memory" (memory 0))
	)
	`)
	require.NoError(err)

	maxUnits := uint64(1)
	cfg, err := NewConfigBuilder().
		WithLimitMaxMemory(1 * MemoryPageSize). // 1 page
		Build()
	require.NoError(err)
	runtime := New(logging.NoLog{}, cfg, nil)
	err = runtime.Initialize(context.Background(), wasm, maxUnits)
	require.NoError(err)

	length, err := runtime.Memory().Len()
	require.NoError(err)
	require.Equal(uint64(0x10000), length)

	// attempt to grow memory to 2 pages which exceeds the limit
	_, err = runtime.Memory().Grow(1)
	require.ErrorContains(err, "failed to grow memory by `1`")
}

func TestWriteExceedsLimitMaxMemory(t *testing.T) {
	require := require.New(t)

	wasm, err := wasmtime.Wat2Wasm(`
	(module
	
	  (memory 1) ;; 1 pages
	  (export "memory" (memory 0))
	)
	`)
	require.NoError(err)

	maxUnits := uint64(1)
	cfg, err := NewConfigBuilder().
		WithLimitMaxMemory(1 * MemoryPageSize). // 1 page
		Build()
	require.NoError(err)
	runtime := New(logging.NoLog{}, cfg, nil)
	err = runtime.Initialize(context.Background(), wasm, maxUnits)
	require.NoError(err)
	maxMemory, err := runtime.Memory().Len()
	require.NoError(err)

	bytes := utils.RandomBytes(int(maxMemory) + 1)
	err = runtime.Memory().Write(0, bytes)
	require.Error(err, "write memory failed: invalid memory size")
}

func TestWithMaxWasmStack(t *testing.T) {
	require := require.New(t)
	wasm, err := wasmtime.Wat2Wasm(`
	(module $test
	(type (;0;) (func (result i32)))
	(export "get_guest" (func 0))
	(func (;0;) (type 0) (result i32)
		(local i32)
		i32.const 1
	  )
	) 
	`)
	require.NoError(err)

	maxUnits := uint64(4)
	cfg, err := NewConfigBuilder().
		WithMaxWasmStack(720).
		Build()
	require.NoError(err)
	runtime := New(logging.NoLog{}, cfg, nil)
	err = runtime.Initialize(context.Background(), wasm, maxUnits)
	require.NoError(err)
	_, err = runtime.Call(context.Background(), "get")
	require.NoError(err)

	// stack is ok for 1 call.
	cfg, err = NewConfigBuilder().
		WithMaxWasmStack(500).
		Build()
	require.NoError(err)
	runtime = New(logging.NoLog{}, cfg, nil)
	err = runtime.Initialize(context.Background(), wasm, maxUnits)
	require.NoError(err)
	// exceed the stack limit
	_, err = runtime.Call(context.Background(), "get")
	require.ErrorIs(err, ErrTrapStackOverflow)
}

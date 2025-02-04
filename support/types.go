// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs types_def.go

package support // import "go.opentelemetry.io/ebpf-profiler/support"

const (
	FrameMarkerUnknown	= 0x0
	FrameMarkerErrorBit	= 0x80
	FrameMarkerPython	= 0x1
	FrameMarkerNative	= 0x3
	FrameMarkerPHP		= 0x2
	FrameMarkerPHPJIT	= 0x9
	FrameMarkerKernel	= 0x4
	FrameMarkerHotSpot	= 0x5
	FrameMarkerRuby		= 0x6
	FrameMarkerPerl		= 0x7
	FrameMarkerV8		= 0x8
	FrameMarkerDotnet	= 0xa
	FrameMarkerAbort	= 0xff
)

const (
	ProgUnwindStop		= 0x0
	ProgUnwindNative	= 0x1
	ProgUnwindHotspot	= 0x2
	ProgUnwindPython	= 0x4
	ProgUnwindPHP		= 0x5
	ProgUnwindRuby		= 0x6
	ProgUnwindPerl		= 0x3
	ProgUnwindV8		= 0x7
	ProgUnwindDotnet	= 0x8
)

const (
	DeltaCommandFlag	= 0x8000

	MergeOpcodeNegative	= 0x80
)

const (
	EventTypeGenericPID = 0x1
)

const MaxFrameUnwinds = 0x80

const (
	MetricIDBeginCumulative = 0x60
)

const (
	BitWidthPID	= 0x20
	BitWidthPage	= 0x40
)

const (
	StackDeltaBucketSmallest	= 0x8
	StackDeltaBucketLargest		= 0x17

	StackDeltaPageBits	= 0x10
	StackDeltaPageMask	= 0xffff
)

const (
	HSTSIDIsStubBit		= 0x3f
	HSTSIDHasFrameBit	= 0x3e
	HSTSIDStackDeltaBit	= 0x38
	HSTSIDStackDeltaMask	= 0x3f
	HSTSIDStackDeltaScale	= 0x8
	HSTSIDSegMapBit		= 0x0
	HSTSIDSegMapMask	= 0xffffffffffffff
)

const (
	PerfMaxStackDepth = 0x7f
)

const (
	TraceOriginUnknown	= 0x0
	TraceOriginSampling	= 0x1
	TraceOriginOffCPU	= 0x2
)

const OffCPUThresholdMax = 0x3e8

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//! Raw protobuf message definitions.

#![allow(missing_docs)]

// Simply include protobuf definitions generated by `build.rs`.
include!(concat!(env!("OUT_DIR"), "/symbfile.rs"));

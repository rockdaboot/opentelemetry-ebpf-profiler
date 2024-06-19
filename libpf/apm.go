/*
 * Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
 * or more contributor license agreements. Licensed under the Apache License 2.0.
 * See the file "LICENSE" for details.
 */

package libpf

type APMSpanID [8]byte
type APMTraceID [16]byte
type APMTransactionID = APMSpanID

var InvalidAPMSpanID = APMSpanID{0, 0, 0, 0, 0, 0, 0, 0}

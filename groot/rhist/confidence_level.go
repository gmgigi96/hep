// Copyright ©2022 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rhist

import (
	"fmt"
	"reflect"

	"go-hep.org/x/hep/groot/rbase"
	"go-hep.org/x/hep/groot/rbytes"
	"go-hep.org/x/hep/groot/root"
	"go-hep.org/x/hep/groot/rtypes"
	"go-hep.org/x/hep/groot/rvers"
)

// ConfidenceLevel holds information about 95% confidence level limits.
type ConfidenceLevel struct {
	base   rbase.Object `groot:"BASE-TObject"` // base class
	fNNMC  int32        `groot:"fNNMC"`
	fDtot  int32        `groot:"fDtot"`
	fStot  float64      `groot:"fStot"`
	fBtot  float64      `groot:"fBtot"`
	fTSD   float64      `groot:"fTSD"`
	fNMC   float64      `groot:"fNMC"`
	fMCL3S float64      `groot:"fMCL3S"`
	fMCL5S float64      `groot:"fMCL5S"`
	fTSB   []float64    `groot:"fTSB,meta=[fNNMC]"`
	fTSS   []float64    `groot:"fTSS,meta=[fNNMC]"`
	fLRS   []float64    `groot:"fLRS,meta=[fNNMC]"`
	fLRB   []float64    `groot:"fLRB,meta=[fNNMC]"`
	fISS   []int32      `groot:"fISS,meta=[fNNMC]"`
	fISB   []int32      `groot:"fISB,meta=[fNNMC]"`
}

func (*ConfidenceLevel) Class() string {
	return "TConfidenceLevel"
}

func (*ConfidenceLevel) RVersion() int16 {
	return rvers.ConfidenceLevel
}

// MarshalROOT implements rbytes.Marshaler
func (o *ConfidenceLevel) MarshalROOT(w *rbytes.WBuffer) (int, error) {
	if w.Err() != nil {
		return 0, w.Err()
	}

	pos := w.WriteVersion(o.RVersion())

	if _, err := o.base.MarshalROOT(w); err != nil {
		return 0, err
	}

	w.WriteI32(int32(o.fNNMC))
	w.WriteI32(o.fDtot)
	w.WriteF64(o.fStot)
	w.WriteF64(o.fBtot)
	w.WriteF64(o.fTSD)
	w.WriteF64(o.fNMC)
	w.WriteF64(o.fMCL3S)
	w.WriteF64(o.fMCL5S)
	w.WriteI8(1) // is-array
	w.WriteFastArrayF64(o.fTSB[:o.fNNMC])
	w.WriteI8(1) // is-array
	w.WriteFastArrayF64(o.fTSS[:o.fNNMC])
	w.WriteI8(1) // is-array
	w.WriteFastArrayF64(o.fLRS[:o.fNNMC])
	w.WriteI8(1) // is-array
	w.WriteFastArrayF64(o.fLRB[:o.fNNMC])
	w.WriteI8(1) // is-array
	w.WriteFastArrayI32(o.fISS[:o.fNNMC])
	w.WriteI8(1) // is-array
	w.WriteFastArrayI32(o.fISB[:o.fNNMC])

	return w.SetByteCount(pos, o.Class())
}

// UnmarshalROOT implements rbytes.Unmarshaler
func (o *ConfidenceLevel) UnmarshalROOT(r *rbytes.RBuffer) error {
	if r.Err() != nil {
		return r.Err()
	}

	start := r.Pos()
	vers, pos, bcnt := r.ReadVersion(o.Class())
	if vers > rvers.ConfidenceLevel {
		panic(fmt.Errorf("rhist: invalid TConfidenceLevel version=%d > %d", vers, rvers.ConfidenceLevel))
	}

	if err := o.base.UnmarshalROOT(r); err != nil {
		return err
	}
	o.fNNMC = r.ReadI32()
	o.fDtot = r.ReadI32()
	o.fStot = r.ReadF64()
	o.fBtot = r.ReadF64()
	o.fTSD = r.ReadF64()
	o.fNMC = r.ReadF64()
	o.fMCL3S = r.ReadF64()
	o.fMCL5S = r.ReadF64()
	_ = r.ReadI8() // is-array
	o.fTSB = rbytes.ResizeF64(nil, int(o.fNNMC))
	r.ReadArrayF64(o.fTSB)
	_ = r.ReadI8() // is-array
	o.fTSS = rbytes.ResizeF64(nil, int(o.fNNMC))
	r.ReadArrayF64(o.fTSS)
	_ = r.ReadI8() // is-array
	o.fLRS = rbytes.ResizeF64(nil, int(o.fNNMC))
	r.ReadArrayF64(o.fLRS)
	_ = r.ReadI8() // is-array
	o.fLRB = rbytes.ResizeF64(nil, int(o.fNNMC))
	r.ReadArrayF64(o.fLRB)
	_ = r.ReadI8() // is-array
	o.fISS = rbytes.ResizeI32(nil, int(o.fNNMC))
	r.ReadArrayI32(o.fISS)
	_ = r.ReadI8() // is-array
	o.fISB = rbytes.ResizeI32(nil, int(o.fNNMC))
	r.ReadArrayI32(o.fISB)

	r.CheckByteCount(pos, bcnt, start, o.Class())
	return r.Err()
}

func init() {
	f := func() reflect.Value {
		var o ConfidenceLevel
		return reflect.ValueOf(&o)
	}
	rtypes.Factory.Add("TConfidenceLevel", f)
}

var (
	_ root.Object        = (*ConfidenceLevel)(nil)
	_ rbytes.RVersioner  = (*ConfidenceLevel)(nil)
	_ rbytes.Marshaler   = (*ConfidenceLevel)(nil)
	_ rbytes.Unmarshaler = (*ConfidenceLevel)(nil)
)

package statistics

// Copyright 2012 G.vd.Schoot. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
// Interface is used throughout the package.
//
type Interface interface {
	Value(int) float64
	SetValue(int, float64)
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

type Float64 []float64

func (f *Float64) Value(i int) float64         { return (*f)[i] }
func (f *Float64) SetValue(i int, val float64) { (*f)[i] = val }
func (f *Float64) Len() int                    { return len(*f) }
func (f *Float64) Less(i, j int) bool          { return (*f)[i] < (*f)[j] }
func (f *Float64) Swap(i, j int)               { (*f)[i], (*f)[j] = (*f)[j], (*f)[i] }

type Int64 []int64

func (f *Int64) Value(i int) float64           { return float64((*f)[i]) }
func (f *Int64) SetValue(i int, value float64) { (*f)[i] = int64(value) }
func (f *Int64) Len() int                      { return len(*f) }
func (f *Int64) Less(i, j int) bool            { return (*f)[i] < (*f)[j] }
func (f *Int64) Swap(i, j int)                 { (*f)[i], (*f)[j] = (*f)[j], (*f)[i] }

//
// Strider strides over the data. It makes steps of size "stride".
//

type Strider struct {
	Interface
	stride int
}

func NewStrider(data Interface, stride int) *Strider {
	return &Strider{data, stride}
}

func (p *Strider) Value(i int) float64 {
	return p.Interface.Value(i * p.stride)
}

func (p *Strider) SetValue(i int, value float64) {
	p.Interface.SetValue(i*p.stride, value)
}

func (p *Strider) Len() int {
	return p.Interface.Len() / p.stride
}

func (p *Strider) Less(i, j int) bool {
	return p.Interface.Less(i*p.stride, j*p.stride)
}

func (p *Strider) Swap(i, j int) {
	p.Interface.Swap(i*p.stride, j*p.stride)
}

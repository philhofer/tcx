package tcx

import (
	"github.com/philhofer/vec"
)

type TrackSpline struct {
	lat  *vec.CubicSplineInterpolation
	long *vec.CubicSplineInterpolation
	spd  *vec.CubicSplineInterpolation
	alt  *vec.CubicSplineInterpolation
}

func (t *TrackSpline) Lat(dist float64) float64 {
	return t.lat.F(dist)
}

func (t *TrackSpline) Long(dist float64) float64 {
	return t.long.F(dist)
}

func (t *TrackSpline) Speed(dist float64) float64 {
	return t.spd.F(dist)
}

func (t *TrackSpline) Alt(dist float64) float64 {
	return t.alt.F(dist)
}

func Spline(trk Track) *TrackSpline {
	Npts := len(trk.Pt)
	lats := make([]float64, Npts)
	longs := make([]float64, Npts)
	spds := make([]float64, Npts)
	alts := make([]float64, Npts)
	dists := make([]float64, Npts)
	for i := 0; i < Npts; i++ {
		lats[i] = trk.Pt[i].Lat
		longs[i] = trk.Pt[i].Long
		spds[i] = trk.Pt[i].Speed
		alts[i] = trk.Pt[i].Alt
		dists[i] = trk.Pt[i].Dist
	}
	latsplin := vec.CubicSpline(vec.MakeBiVariateData(dists, lats))
	longssplin := vec.CubicSpline(vec.MakeBiVariateData(dists, longs))
	spdssplin := vec.CubicSpline(vec.MakeBiVariateData(dists, spds))
	altssplin := vec.CubicSpline(vec.MakeBiVariateData(dists, alts))
	return &TrackSpline{lat: latsplin, long: longssplin, spd: spdssplin, alt: altssplin}
}

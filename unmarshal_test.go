package tcx

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	db, err := ReadFile("files/sample.tcx")
	if err != nil {
		t.Error(err)
	}
	npts := len(db.Acts.Act[0].Laps[0].Trk.Pt)
	nlaps := len(db.Acts.Act[0].Laps)
	nacts := len(db.Acts.Act)
	if nlaps != 1 || nacts != 1 {
		t.Error("Incorrectly parsed # laps or # acts not equal to 1.")
		t.Error("# Laps parsed:", nlaps)
		t.Error("# Activities parsed:", nacts)
	}
	finalPt := db.Acts.Act[0].Laps[0].Trk.Pt[npts-1]
	if finalPt.Lat != 37.8765614 || finalPt.Long != -122.4601646 {
		t.Error("Lat/Long parsed incorrectly.")
		t.Error("Expected lat/long of 37.8765614/-122.4601646")
		t.Error("Got:", finalPt.Lat, finalPt.Long)
	}
	return
}

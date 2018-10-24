package bench

import (
	"github.com/bsm/vast"
)

var creatives []vast.Creative // we use global variables in benchmarks to prevent the optimizer to remove our useless operations

func initCreatives() {
	creatives = []vast.Creative{}
	for i := 0; i < 5; i++ {
		creatives = append(creatives, vast.Creative{
			Linear: &vast.Linear{
				SkipOffset: &vast.Offset{},
				Duration: &[]vast.Duration{0}[0],
				Icons: []vast.Icon{},
				TrackingEvents: []vast.Tracking{
					vast.Tracking {
						Offset: &vast.Offset{},
						URI: vast.URI("uri"),
					},
				},
				MediaFiles: []vast.MediaFile{
				},
				CreativeExtensions: &vast.CreativeExtensions{
				},
			},
			CompanionAds: &vast.CompanionAds{
				Companions: []vast.Companion{
				},
			},
		})
	}
}

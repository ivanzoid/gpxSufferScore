Given GPX file, calculates "suffer score" = heart rate / speed. Output in CSV file (you can use gnuplot for viewing result).

NOTE: you need to add the following line to type Wpt in gpx.go in github.com/ptrv/go-gpx:

Hr           int     `xml:"extensions>TrackPointExtension>hr,omitempty"`
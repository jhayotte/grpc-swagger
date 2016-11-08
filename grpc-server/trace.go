package main

import (
	"time"

	pb "../pointpb"
)

func GetPosition(assetId string) pb.Point {
	// Simulate remote database overhead

	sq := pb.Point{
		AssetID:  "VehicleX",
		Lat:      1.22215,
		Lng:      1.223,
		Datetime: time.Now().String(),
	}

	return sq
}

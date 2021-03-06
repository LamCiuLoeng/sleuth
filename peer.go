// Copyright 2016 Afshin Darian. All rights reserved.
// Use of this source code is governed by The MIT License
// that can be found in the LICENSE file.

package sleuth

// peer describes the location of a peer on the sleuth network and the service,
// if any, that it offers.
type peer struct {
	// name is the short public name attached to all events/peers.
	name string
	// node is the full peer node name used for whispering.
	node string
	// service is the name of the service being offered by a peer.
	service string
	// version is the optional service version running on a peer.
	version string
}

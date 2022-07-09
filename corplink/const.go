package corplink

// corplink version of device.WGIdentifier
const WGIdentifier = "CorpLink v1 vpn@feilian-----------"

const (
	// protocol version to support modified identifier
	EnvKeyProtocolVersion = "CORPLINK_PROTOCOL_VERSION"
	// for wg over tcp support, will implement in the future
	EnvKeyNetworkType = "CORPLINK_NETWORK_TYPE"
)

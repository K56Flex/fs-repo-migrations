package multiaddr

/*
Multiaddr is a cross-protocol, cross-platform format for representing
internet addresses. It emphasizes explicitness and self-description.
Learn more here: https://github.com/jbenet/multiaddr

Multiaddrs have both a binary and string representation.

    import ma "github.com/jbenet/go-multiaddr"

    addr, err := ma.NewMultiaddr("/ip4/1.2.3.4/tcp/80")
    // err non-nil when parsing failed.

*/
type Multiaddr interface {
	// Bytes returns the []byte representation of this Multiaddr
	Bytes() []byte

	// String returns the string representation of this Multiaddr
	// (may panic if internal state is corrupted)
	String() string

	// Protocols returns the list of Protocols this Multiaddr includes
	// will panic if protocol code incorrect (and bytes accessed incorrectly)
	Protocols() []Protocol

	// ValueForProtocol returns the value (if any) following the specified protocol
	ValueForProtocol(code int) (string, error)
}

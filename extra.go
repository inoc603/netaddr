package netaddr

func (ip IP) High() uint64 {
	return ip.addr.hi
}

func (ip IP) Low() uint64 {
	return ip.addr.lo
}

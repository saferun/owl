/*
 * Copyright 2019-2020 by Nedim Sabic Sabic
 * https://www.fibratus.io
 * All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ip

import (
	"net"

	"github.com/mel2oo/win32/ntdll/ip2string"
)

// ToIPv4 accepts an integer IP address in network byte order and returns an IP-typed address.
func ToIPv4(ip uint32) net.IP {
	return net.IPv4(byte(ip), byte(ip>>8), byte(ip>>16), byte(ip>>24))
}

// ToIPv6 converts the buffer with IPv6 address in network byte order to an IP-typed address.
func ToIPv6(buffer []byte) net.IP {
	ipv6 := make([]uint16, 46)
	return net.ParseIP(ip2string.RtlIpv6AddressToString(buffer, ipv6))
}

package stats

import (
	"github.com/shirou/gopsutil/v3/net"
)

type NetworkData struct {
	TcpConnections   []net.ConnectionStat    `json:"tcp_connections,omitempty"`
	UdpConnections   []net.ConnectionStat    `json:"udp_connections,omitempty"`
	FilterCounters   []net.FilterStat        `json:"filter_counters,omitempty"`
	IoCounters       []net.IOCountersStat    `json:"io_counters,omitempty"`
	Interfaces       net.InterfaceStatList   `json:"interfaces,omitempty"`
	ProtocolCounters []net.ProtoCountersStat `json:"protocol_counters,omitempty"`
}

var (
	tcpConnections []net.ConnectionStat
	udpConnections []net.ConnectionStat
	filterCounters []net.FilterStat
	ioCounters     []net.IOCountersStat
	interfaces     net.InterfaceStatList
	protoCounters  []net.ProtoCountersStat
)

func GetNetworkData(filterData, interfaceData, iostatData, protocolData bool) (networkData NetworkData, err error) {
	if tcpConnections, err = net.Connections("tcp"); err != nil {
		return NetworkData{}, err
	}

	if udpConnections, err = net.Connections("tcp"); err != nil {
		return NetworkData{}, err
	}

	if filterData {
		if filterCounters, err = net.FilterCounters(); err != nil {
			return NetworkData{}, err
		}
	}

	if iostatData {
		if ioCounters, err = net.IOCounters(true); err != nil {
			return NetworkData{}, err
		}
	}

	if interfaceData {
		if interfaces, err = net.Interfaces(); err != nil {
			return NetworkData{}, err
		}
	}

	if protocolData {
		if protoCounters, err = net.ProtoCounters([]string{"ip", "icmp", "icmpmsg", "tcp", "udp", "udplite"}); err != nil {
			return NetworkData{}, err
		}
	}

	networkData = NetworkData{
		TcpConnections:   tcpConnections,
		UdpConnections:   udpConnections,
		FilterCounters:   filterCounters,
		IoCounters:       ioCounters,
		Interfaces:       interfaces,
		ProtocolCounters: protoCounters,
	}

	return networkData, nil
}

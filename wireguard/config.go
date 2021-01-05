package wireguard

type Config struct {
	// Wireguard configuration
	Enabled                    bool
	RouteSource                string
	ListeningPort              int
	MarkDoNotRouteViaWireguard int
	RoutingRulePriority        int
	WorkloadRoutingTableIndex  int
	InterfaceName              string
	MTU                        int
}

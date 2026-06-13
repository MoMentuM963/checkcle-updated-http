package pocketbase

import "time"

type AuthResponse struct {
	Token  string      `json:"token"`
	Record interface{} `json:"record"`
}

type MetricsRecord struct {
	ServiceName       string `json:"service_name"`
	Host              string `json:"host"`
	Uptime            float64 `json:"uptime"`
	ResponseTime      int64   `json:"response_time"`
	LastChecked       string  `json:"last_checked"`
	Port              int     `json:"port,omitempty"`
	Domain            string  `json:"domain,omitempty"`
	HeartbeatInterval int     `json:"heartbeat_interval,omitempty"`
	MaxRetries        int     `json:"max_retries,omitempty"`
	NotificationID    string  `json:"notification_id,omitempty"`
	TemplateID        string  `json:"template_id,omitempty"`
	ServiceType       string  `json:"service_type"`
	Status            string  `json:"status"`
	URL               string  `json:"url,omitempty"`
	Alerts            string  `json:"alerts,omitempty"`
	StatusCodes       string  `json:"status_codes,omitempty"`
	Keyword           string  `json:"keyword,omitempty"`
	ErrorMessage      string  `json:"error_message,omitempty"`
	Details           string  `json:"details,omitempty"`
	CheckedAt         string  `json:"checked_at"`
}

type PingDataRecord struct {
	ServiceID     string    `json:"service_id"`
	Timestamp     time.Time `json:"timestamp"`
	ResponseTime  int64     `json:"response_time"`
	Status        string    `json:"status"`
	PacketLoss    string    `json:"packet_loss"`
	Latency       string    `json:"latency"`
	MaxRTT        string    `json:"max_rtt"`
	MinRTT        string    `json:"min_rtt"`
	PacketsSent   string    `json:"packets_sent"`
	PacketsRecv   string    `json:"packets_recv"`
	AvgRTT        string    `json:"avg_rtt"`
	RTTs          string    `json:"rtts"`
	Details       string    `json:"details,omitempty"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	RegionName    string    `json:"region_name,omitempty"`
	AgentID       string    `json:"agent_id,omitempty"`
}

type UptimeDataRecord struct {
	ServiceID     string    `json:"service_id"`
	Timestamp     time.Time `json:"timestamp"`
	ResponseTime  int64     `json:"response_time"`
	Status        string    `json:"status"`
	Packets       string    `json:"packets"`
	Latency       string    `json:"latency"`
	StatusCodes   string    `json:"status_codes"`
	Keyword       string    `json:"keyword"`
	ErrorMessage  string    `json:"error_message"`
	Details       string    `json:"details"`
	Region        string    `json:"region,omitempty"`
	RegionID      string    `json:"region_id,omitempty"`
	RegionName    string    `json:"region_name,omitempty"`
	AgentID       string    `json:"agent_id,omitempty"`
}

type DNSDataRecord struct {
	ServiceID    string    `json:"service_id"`
	Timestamp    time.Time `json:"timestamp"`
	ResponseTime int64     `json:"response_time"`
	Status       string    `json:"status"`
	QueryType    string    `json:"query_type"`
	ResolveIP    string    `json:"resolve_ip"`
	MsgSize      string    `json:"msg_size"`
	Question     string    `json:"question"`
	Answer       string    `json:"answer"`
	Authority    string    `json:"authority"`
	ErrorMessage string    `json:"error_message,omitempty"`
	Details      string    `json:"details,omitempty"`
	RegionName   string    `json:"region_name,omitempty"`
	AgentID      string    `json:"agent_id,omitempty"`
}

type TCPDataRecord struct {
	ServiceID    string    `json:"service_id"`
	Timestamp    time.Time `json:"timestamp"`
	ResponseTime int64     `json:"response_time"`
	Status       string    `json:"status"`
	Connection   string    `json:"connection"`
	Latency      string    `json:"latency"`
	Port         string    `json:"port"`
	ErrorMessage string    `json:"error_message,omitempty"`
	Details      string    `json:"details,omitempty"`
	RegionName   string    `json:"region_name,omitempty"`
	AgentID      string    `json:"agent_id,omitempty"`
}

// SSL Data Record remains unchanged - no regional agent fields
type SSLDataRecord struct {
	ServiceID     string    `json:"service_id"`
	Timestamp     time.Time `json:"timestamp"`
	ResponseTime  int64     `json:"response_time"`
	Status        string    `json:"status"`
	ValidFrom     string    `json:"valid_from"`
	ValidTill     string    `json:"valid_till"`
	DaysLeft      int       `json:"days_left"`
	Issuer        string    `json:"issuer"`
	Subject       string    `json:"subject"`
	SerialNumber  string    `json:"serial_number"`
	Algorithm     string    `json:"algorithm"`
	SANs          string    `json:"sans"`
	ResolvedIP    string    `json:"resolved_ip"`
	ErrorMessage  string    `json:"error_message,omitempty"`
	Details       string    `json:"details,omitempty"`
}

// Regional Service record structure
type RegionalService struct {
	ID               string `json:"id"`
	RegionName       string `json:"region_name"`
	Status           string `json:"status"`
	AgentID          string `json:"agent_id"`
	AgentIPAddress   string `json:"agent_ip_address"`
	Connection       string `json:"connection"`
	Token            string `json:"token"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
}

type RegionalServicesResponse struct {
	Page       int               `json:"page"`
	PerPage    int               `json:"perPage"`
	TotalItems int               `json:"totalItems"`
	TotalPages int               `json:"totalPages"`
	Items      []RegionalService `json:"items"`
}

type Service struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	Host               string    `json:"host"`
	Uptime             float64   `json:"uptime"`
	ResponseTime       int64     `json:"response_time"`
	LastChecked        string    `json:"last_checked"`
	Port               int       `json:"port"`
	Domain             string    `json:"domain"`
	HeartbeatInterval  int       `json:"heartbeat_interval"`
	MaxRetries         int       `json:"max_retries"`
	NotificationID     string    `json:"notification_id"`
	TemplateID         string    `json:"template_id"`
	ServiceType        string    `json:"service_type"`
	Status             string    `json:"status"`
	URL                string    `json:"url"`
	Alerts             string    `json:"alerts"`
	StatusCodes        string    `json:"status_codes"`
	Keyword            string    `json:"keyword"`	
	Method             string    `json:"method"`
	Headers            string    `json:"headers"`
	Body               string    `json:"body"`
	Created            string    `json:"created"`
	Updated            string    `json:"updated"`
}

type ServicesResponse struct {
	Page       int       `json:"page"`
	PerPage    int       `json:"perPage"`
	TotalItems int       `json:"totalItems"`
	TotalPages int       `json:"totalPages"`
	Items      []Service `json:"items"`
}
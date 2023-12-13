package main

type DockerCompose struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]Service `yaml:"services,omitempty"`
	Networks map[string]Network `yaml:"networks,omitempty"`
	Volumes  map[string]Volume  `yaml:"volumes,omitempty"`
	Secrets  map[string]Secret  `yaml:"secrets,omitempty"`
	Configs  map[string]Config  `yaml:"configs,omitempty"`
}

type Service struct {
	Image           string                 `yaml:"image,omitempty"`
	Build           string                 `yaml:"build,omitempty"`
	NetworkMode     string                 `yaml:"network_mode,omitempty"`
	Networks        map[string]interface{} `yaml:"networks,omitempty"`
	Ports           []string               `yaml:"ports,omitempty"`
	Expose          []string               `yaml:"expose,omitempty"`
	Environment     []string               `yaml:"environment,omitempty"`
	EnvFile         []string               `yaml:"env_file,omitempty"`
	DependsOn       []string               `yaml:"depends_on,omitempty"`
	Command         string                 `yaml:"command,omitempty"`
	Entrypoint      []string               `yaml:"entrypoint,omitempty"`
	Labels          map[string]string      `yaml:"labels,omitempty"`
	WorkingDir      string                 `yaml:"working_dir,omitempty"`
	Volumes         []string               `yaml:"volumes,omitempty"`
	VolumesFrom     []string               `yaml:"volumes_from,omitempty"`
	CapAdd          []string               `yaml:"cap_add,omitempty"`
	CapDrop         []string               `yaml:"cap_drop,omitempty"`
	Devices         []string               `yaml:"devices,omitempty"`
	Links           []string               `yaml:"links,omitempty"`
	ExternalLinks   []string               `yaml:"external_links,omitempty"`
	ExtraHosts      []string               `yaml:"extra_hosts,omitempty"`
	Ulimits         map[string]Ulimit      `yaml:"ulimits,omitempty"`
	HealthCheck     HealthCheck            `yaml:"healthcheck,omitempty"`
	Sysctls         map[string]string      `yaml:"sysctls,omitempty"`
	SecurityOpt     []string               `yaml:"security_opt,omitempty"`
	Logging         Logging                `yaml:"logging,omitempty"`
	Tmpfs           []string               `yaml:"tmpfs,omitempty"`
	StopGracePeriod string                 `yaml:"stop_grace_period,omitempty"`
	StopSignal      string                 `yaml:"stop_signal,omitempty"`
	Tty             bool                   `yaml:"tty,omitempty"`
	Pid             string                 `yaml:"pid,omitempty"`
	ContainerName   string                 `yaml:"container_name,omitempty"`
	User            string                 `yaml:"user,omitempty"`
	Dns             []string               `yaml:"dns,omitempty"`
	Restart string `yaml:"restart,omitempty"`
}

type Network struct {
	Driver             string            `yaml:"driver,omitempty"`
	IPAM               IPAM              `yaml:"ipam,omitempty"`
	Internal           bool              `yaml:"internal,omitempty"`
	External           bool              `yaml:"external,omitempty"`
	Labels             map[string]string `yaml:"labels,omitempty"`
	Attachable         bool              `yaml:"attachable,omitempty"`
	Containers         map[string]string `yaml:"containers,omitempty"`
	Name               string            `yaml:"name,omitempty"`
	EnableIPv6         bool              `yaml:"enable_ipv6,omitempty"`
	IPPrefixLength     int               `yaml:"ip_prefix_length,omitempty"`
	Gateway            string            `yaml:"gateway,omitempty"`
	IPAMDriver         string            `yaml:"ipam_driver,omitempty"`
	IPAMOpt            map[string]string `yaml:"ipam_opt,omitempty"`
	InternalName       string            `yaml:"internal_name,omitempty"`
	ExternalName       string            `yaml:"external_name,omitempty"`
	AttachableName     string            `yaml:"attachable_name,omitempty"`
	EnableIPv6Name     string            `yaml:"enable_ipv6_name,omitempty"`
	IPPrefixLengthName int               `yaml:"ip_prefix_length_name,omitempty"`
	GatewayName        string            `yaml:"gateway_name,omitempty"`
}

type IPAM struct {
	Driver  string          `yaml:"driver,omitempty"`
	Config  []IPAMConfig    `yaml:"config,omitempty"`
	Options map[string]bool `yaml:"options,omitempty"`
}

type IPAMConfig struct {
	Subnet             string            `yaml:"subnet,omitempty"`
	IPRange            string            `yaml:"ip_range,omitempty"`
	Gateway            string            `yaml:"gateway,omitempty"`
	AuxiliaryAddresses map[string]string `yaml:"auxiliary_addresses,omitempty"`
}

type Volume struct {
	Driver           string            `yaml:"driver,omitempty"`
	DriverOpts       map[string]string `yaml:"driver_opts,omitempty"`
	External         map[string]bool   `yaml:"external,omitempty"`
	Labels           map[string]string `yaml:"labels,omitempty"`
	Name             string            `yaml:"name,omitempty"`
	Attachable       bool              `yaml:"attachable,omitempty"`
	DriverName       string            `yaml:"driver_name,omitempty"`
	ExternalName     string            `yaml:"external_name,omitempty"`
	AttachableName   string            `yaml:"attachable_name,omitempty"`
	ExternalNameName string            `yaml:"external_name_name,omitempty"`
}

type Secret struct {
	File         string `yaml:"file,omitempty"`
	External     bool   `yaml:"external,omitempty"`
	Name         string `yaml:"name,omitempty"`
	ExternalName string `yaml:"external_name,omitempty"`
}

type Config struct {
	File         string `yaml:"file,omitempty"`
	External     bool   `yaml:"external,omitempty"`
	Name         string `yaml:"name,omitempty"`
	ExternalName string `yaml:"external_name,omitempty"`
}

type Ulimit struct {
	Name string `yaml:"name,omitempty"`
	Soft int    `yaml:"soft,omitempty"`
	Hard int    `yaml:"hard,omitempty"`
}

type HealthCheck struct {
	Test        []string `yaml:"test,omitempty"`
	Interval    string   `yaml:"interval,omitempty"`
	Timeout     string   `yaml:"timeout,omitempty"`
	Retries     int      `yaml:"retries,omitempty"`
	StartPeriod string   `yaml:"start_period,omitempty"`
}

type Logging struct {
	Driver  string            `yaml:"driver,omitempty"`
	Options map[string]string `yaml:"options,omitempty"`
}

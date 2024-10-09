package agent

import (
	"github.com/pelletier/go-toml/v2"
)

const (
	ConfigFilePath       = "/run/peerpod/agent-config.toml"
	ServerAddr           = "unix:///run/kata-containers/agent.sock"
	GuestComponentsProcs = "none"
	EnableSignatureVerification = "true"
	ImagePolicyFile = "kbs:///default/security-policy/coco-caa"
)

type agentConfig struct {
	ServerAddr           string `toml:"server_addr"`
	GuestComponentsProcs string `toml:"guest_components_procs"`
	ImageRegistryAuth    string `toml:"image_registry_auth,omitempty"`
	EnableSignatureVerification string `toml:"enable_signature_verification"`
	ImagePolicyFile string `toml:"image_policy_file"`
}

func CreateConfigFile(authJsonPath string) (string, error) {
	var imageRegistryAuth string
	if authJsonPath != "" {
		imageRegistryAuth = "file://" + authJsonPath
	}

	config := agentConfig{
		ServerAddr:           ServerAddr,
		GuestComponentsProcs: GuestComponentsProcs,
		ImageRegistryAuth:    imageRegistryAuth,
		EnableSignatureVerification: EnableSignatureVerification,
		ImagePolicyFile: ImagePolicyFile,
	}

	bytes, err := toml.Marshal(config)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

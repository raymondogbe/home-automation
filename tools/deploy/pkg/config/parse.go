package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/jakewright/home-automation/libraries/go/oops"
	"github.com/jakewright/home-automation/tools/deploy/pkg/output"
)

type config struct {
	Repository string              `yaml:"repository"`
	Targets    map[string]*target  `yaml:"targets"`
	Services   map[string]*service `yaml:"services"`
}

type target struct {
	// Common
	Host   string `yaml:"host"`
	System string `yaml:"system"`

	// Systemd
	Username     string `yaml:"username"`
	Directory    string `yaml:"directory"`
	Architecture string `yaml:"architecture"`

	// Kubernetes
	KubeContext      string `yaml:"kube_context"`
	Namespace        string `yaml:"namespace"`
	DockerRegistry   string `yaml:"docker_registry"`
	DockerRepository string `yaml:"docker_repository"`
}

type dockerConfig struct {
	Dockerfile string            `yaml:"dockerfile"`
	Args       map[string]string `yaml:"args"`
}

type service struct {
	TargetNames []string      `yaml:"targets"`
	Language    string        `yaml:"language"`
	EnvFiles    []string      `yaml:"env_files"`
	Docker      *dockerConfig `yaml:"docker"`
}

// Init reads and validates config
func Init(filename string) (err error) {
	op := output.Info("Reading config from %v", filename)
	defer func() {
		if err == nil {
			op.Success()
		} else {
			op.Failed()
		}
	}()

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return oops.WithMessage(err, "failed to read config file")
	}

	raw := &config{}
	if err := yaml.Unmarshal(b, &raw); err != nil {
		return oops.WithMessage(err, "failed to unmarshal config")
	}

	targets := make(map[string]*Target)

	for name, t := range raw.Targets {
		switch t.System {
		case SysDocker, SysKubernetes, SysSystemd: // ok
		default:
			return oops.InternalService("Invalid system %q for target %q", t.System, name)
		}

		switch t.Architecture {
		case "", ArchARMv6: // ok
		default:
			return oops.InternalService("Invalid architecture %q for target %q", t.Architecture, name)
		}

		targets[name] = &Target{
			name:             name,
			host:             t.Host,
			system:           t.System,
			username:         t.Username,
			directory:        t.Directory,
			architecture:     t.Architecture,
			kubeContext:      t.KubeContext,
			namespace:        t.Namespace,
			dockerRegistry:   t.DockerRegistry,
			dockerRepository: t.DockerRepository,
		}
	}

	services := make(map[string]*Service)

	for name, s := range raw.Services {
		switch s.Language {
		case LangGo: // ok
		default:
			return oops.InternalService("Invalid language '%s' for service '%s'", s.Language, name)
		}

		var serviceTargets []*Target

		for _, targetName := range s.TargetNames {
			target := targets[targetName]
			if target == nil {
				return oops.InternalService("Invalid target %q for service %q", targetName, name)
			}

			serviceTargets = append(serviceTargets, target)
		}

		services[name] = &Service{
			name:        name,
			targetNames: s.TargetNames,
			targets:     serviceTargets,
			language:    s.Language,
			envFiles:    s.EnvFiles,
			docker: &DockerConfig{
				dockerfile: s.Docker.Dockerfile,
				args:       s.Docker.Args,
			},
		}
	}

	cfg = &Config{
		Repository: raw.Repository,
		Targets:    targets,
		Services:   services,
	}

	return nil
}
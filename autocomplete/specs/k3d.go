// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["k3d"] = model.Subcommand{
		Name:        []string{"k3d"},
		Description: `K3d is a lightweight wrapper to run k3s in Docker`,
		Options: []model.Option{{
			Name:         []string{"-h", "--help"},
			Description:  `Help for k3d`,
			IsPersistent: true,
		}, {
			Name:         []string{"--verbose"},
			Description:  `Verbose output`,
			IsPersistent: true,
		}, {
			Name:         []string{"--trace"},
			Description:  `Trace output`,
			IsPersistent: true,
		}, {
			Name:        []string{"--version"},
			Description: `Print version information`,
		}, {
			Name:         []string{"--timestamps"},
			Description:  `Print timestamp in log output`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"cluster"},
			Description: `Manage k3s clusters`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new k3s cluster`,
				Args: []model.Arg{{
					Name:        "cluster name",
					Description: `Name of the cluster to create`,
				}},
				Options: []model.Option{{
					Name:        []string{"-a", "--agents"},
					Description: `Number of agents to create`,
					Args: []model.Arg{{
						Name:        "number of agents",
						Description: `Number of agents to create`,
					}},
				}, {
					Name:        []string{"--agents-memory"},
					Description: `Memory limit imposed on the agents nodes`,
					Args: []model.Arg{{
						Name:        "memory limit",
						Description: `Memory limit imposed on the agents nodes`,
					}},
				}, {
					Name:        []string{"--api-port"},
					Description: `Specify the Kubernetes API server port exposed on the LoadBalancer`,
					Args: []model.Arg{{
						Name:        "port",
						Description: `Specify the Kubernetes API server port exposed on the LoadBalancer`,
					}},
				}, {
					Name:        []string{"-c", "--config"},
					Description: `Path of a config file to use`,
					Args: []model.Arg{{
						Name:        "path",
						Description: `Path of a config file to use`,
					}},
				}, {
					Name:        []string{"-e", "--env"},
					Description: `Add environment variables to nodes`,
					Args: []model.Arg{{
						Name:        "environment variables",
						Description: `Add environment variables to nodes`,
					}},
				}, {
					Name:        []string{"--gpus"},
					Description: `GPU devices to add to the cluster node containers`,
					Args: []model.Arg{{
						Name:        "devices",
						Description: `GPU devices to add to the cluster node containers`,
					}},
				}, {
					Name:        []string{"-h", "--help"},
					Description: `Help for create`,
				}, {
					Name:        []string{"--host-alias"},
					Description: `Add ip:host[,host,...] mappings`,
					Args: []model.Arg{{
						Name:        "ip:host[,host,...]",
						Description: `Add ip:host[,host,...] mappings`,
					}},
				}, {
					Name:        []string{"--host-pid-mode"},
					Description: `Enable host pid mode of server(s) and agent(s)`,
				}, {
					Name:        []string{"-i", "--image"},
					Description: `Specify k3s image that you want to use for the nodes`,
					Args: []model.Arg{{
						Name:        "image",
						Description: `Specify k3s image that you want to use for the nodes`,
					}},
				}, {
					Name:        []string{"--k3s-arg"},
					Description: `Additional args passed to k3s command`,
					Args: []model.Arg{{
						Name:        "ARG@NODEFILTER[;@NODEFILTER]",
						Description: `Additional args passed to k3s command`,
					}},
				}, {
					Name:        []string{"--k3s-node-label"},
					Description: `Add label to k3s node`,
					Args: []model.Arg{{
						Name:        "KEY[=VALUE][@NODEFILTER[;NODEFILTER...]]",
						Description: `Add label to k3s node`,
					}},
				}, {
					Name:        []string{"--kubeconfig-switch-context"},
					Description: `Directly switch the default kubeconfig's current-context to the new cluster's context (requires --kubeconfig-update-default) (default true)`,
				}, {
					Name:        []string{"--kubeconfig-update-default"},
					Description: `Directly update the default kubeconfig with the new cluster's context (default true)`,
				}, {
					Name:        []string{"--lb-config-override"},
					Description: `Use dotted YAML path syntax to override nginx loadbalancer settings`,
					Args: []model.Arg{{
						Name: "path",
					}},
				}, {
					Name:        []string{"--network"},
					Description: `Specify the docker network to use`,
					Args: []model.Arg{{
						Name:        "network",
						Description: `Specify the docker network to use`,
					}},
				}, {
					Name:        []string{"--no-image-volume"},
					Description: `Don't create a volume for the container's root filesystem`,
				}, {
					Name:        []string{"--no-lb"},
					Description: `Don't create a loadbalancer for the cluster`,
				}, {
					Name:        []string{"--no-rollback"},
					Description: `Don't rollback changes if cluster creation failed`,
				}, {
					Name:        []string{"-p", "--port"},
					Description: `Map ports from the node containers (via the serverlb) to the host`,
					Args: []model.Arg{{
						Name:        "port",
						Description: `Map ports from the node containers (via the serverlb) to the host`,
					}},
				}, {
					Name:        []string{"--registry-config"},
					Description: `Specify path to an extra registries.yaml file`,
					Args: []model.Arg{{
						Name:        "path",
						Description: `Specify path to an extra registries.yaml file`,
					}},
				}, {
					Name:        []string{"--registry-create"},
					Description: `Create a k3d-managed registry and connect it to the cluster (Format: NAME[:HOST][:HOSTPORT] - Example: "k3d cluster create --registry-create mycluster-registry`,
					Args: []model.Arg{{
						Name:        "NAME[:HOST][:HOSTPORT]",
						Description: `Create a k3d-managed registry and connect it to the cluster (Format: NAME[:HOST][:HOSTPORT] - Example: "k3d cluster create --registry-create mycluster-registry`,
					}},
				}, {
					Name:        []string{"--registry-use"},
					Description: `Connect to one or more k3d-managed registries running locally`,
					Args: []model.Arg{{
						Name: "NAME[:HOST][:HOSTPORT]",
					}},
				}, {
					Name:        []string{"--runtime-label"},
					Description: `Add label to container runtime (Format: KEY[=VALUE][@NODEFILTER[;NODEFILTER...]] - Example: "k3d cluster create --agents 2 --runtime-label "my.label@agent:0,1" --runtime-label "other.label=somevalue@server:0""`,
				}, {
					Name:        []string{"-s", "--servers"},
					Description: `Specify how many servers you want to create`,
					Args: []model.Arg{{
						Name:        "int",
						Description: `Specify how many servers you want to create`,
					}},
				}, {
					Name:        []string{"--servers-memory"},
					Description: `Memory limit imposed on the server nodes [From docker]`,
					Args: []model.Arg{{
						Name:        "string",
						Description: `Memory limit imposed on the server nodes [From docker]`,
					}},
				}, {
					Name:        []string{"--subnet"},
					Description: `[Experimental: IPAM] Define a subnet for the newly created container network`,
					Args: []model.Arg{{
						Name: "string",
					}},
				}, {
					Name:        []string{"--timeout"},
					Description: `Rollback changes if cluster couldn't be created in specified duration`,
					Args: []model.Arg{{
						Name:        "duration",
						Description: `Rollback changes if cluster couldn't be created in specified duration`,
					}},
				}, {
					Name:        []string{"--token"},
					Description: `Specify a cluster token. By default, we generate one`,
					Args: []model.Arg{{
						Name:        "string",
						Description: `Specify a cluster token. By default, we generate one`,
					}},
				}, {
					Name:        []string{"-v", "--volume"},
					Description: `Mount volumes into the nodes`,
					Args: []model.Arg{{
						Name:        "[SOURCE:]DEST[@NODEFILTER[;NODEFILTER...]]",
						Description: `Mount volumes into the nodes`,
					}},
				}, {
					Name:        []string{"--wait"},
					Description: `Wait for the server(s) to be ready before returning`,
				}},
			}, {
				Name:        []string{"delete", "del", "rm"},
				Description: `Delete a cluster`,
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to delete`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-a", "--all"},
					Description: `Delete all clusters`,
				}, {
					Name:        []string{"-c", "--config"},
					Description: `Path of a config file to use`,
					Args: []model.Arg{{
						Name:        "path",
						Description: `Path of a config file to use`,
					}},
				}},
			}, {
				Name:        []string{"edit", "update"},
				Description: `Edit a cluster`,
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to edit`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"--port-add"},
					Description: `Map ports from the node containers`,
				}},
			}, {
				Name:        []string{"list", "ls", "get"},
				Description: `List clusters`,
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to list`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"--no-headers"},
					Description: `Don't print headers`,
				}, {
					Name:        []string{"-o", "--output"},
					Description: `Output format`,
					Args: []model.Arg{{
						Name:        "string",
						Description: `Output format`,
						Suggestions: []model.Suggestion{{
							Name: []string{`json`},
						}, {
							Name: []string{`yaml`},
						}},
					}},
				}, {
					Name:        []string{"--token"},
					Description: `Print k3s cluster token`,
				}},
			}, {
				Name:        []string{"start"},
				Description: `Start a cluster`,
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to start`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"a", "--all"},
					Description: `Start all clusters`,
				}, {
					Name:        []string{"--timeout"},
					Description: `Maximum waiting time for '--wait' before canceling/returning`,
					Args: []model.Arg{{
						Name:        "duration",
						Description: `Maximum waiting time for '--wait' before canceling/returning`,
					}},
				}, {
					Name:        []string{"--wait"},
					Description: `Wait for the server(s) (and loadbalancer) to be ready before returning`,
				}},
			}, {
				Name:        []string{"stop"},
				Description: `Stop a cluster`,
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to stop`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-a", "--all"},
					Description: `Stop all clusters`,
				}},
			}},
		}, {
			Name:        []string{"completion"},
			Description: `Generate shell completion scripts`,
			Args: []model.Arg{{
				Name:        "shell",
				Description: `Shell to generate completion script for`,
				Suggestions: []model.Suggestion{{
					Name:        []string{`bash`},
					Description: `Bash shell`,
				}, {
					Name:        []string{`zsh`},
					Description: `Zsh shell`,
				}, {
					Name:        []string{`fish`},
					Description: `Fish shell`,
				}, {
					Name:        []string{`powershell`},
					Description: `Powershell`,
				}},
			}},
		}, {
			Name:        []string{"config"},
			Description: `Manage k3d config`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"init", "create"},
				Description: `Create a new config file`,
				Options: []model.Option{{
					Name:        []string{"-f", "--force"},
					Description: `Force overwrite of target file`,
				}, {
					Name:        []string{"-o", "--output"},
					Description: `Write a default k3d config (default "k3d-default.yaml")`,
				}},
			}, {
				Name:        []string{"migrate", "update"},
				Description: `Migrate a config file to the latest version`,
			}},
		}, {
			Name:        []string{"image"},
			Description: `Manage k3d images`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"import", "load"},
				Description: `Import image(s) from docker into k3d cluster(s)`,
				Args: []model.Arg{{
					Name:        "image",
					Description: `Image to import`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-c", "--cluster"},
					Description: `Cluster to import image(s) into`,
					Args: []model.Arg{{
						Name:        "cluster",
						Description: `Cluster to import image(s) into`,
						Generator:   nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"-k", "--keep-tarball"},
					Description: `Do not delete the tarball containing the saved images from the shared volume`,
				}, {
					Name:        []string{"-t", "--keep-tools"},
					Description: `Do not delete the tools node after import`,
				}, {
					Name:        []string{"-m", "--mode"},
					Description: `Import mode`,
					Args: []model.Arg{{
						Name:        "mode",
						Description: `Import mode`,
						Suggestions: []model.Suggestion{{Name: []string{`auto`}}, {Name: []string{`direct`}}, {Name: []string{`tools`}}},
					}},
				}},
			}},
		}, {
			Name:        []string{"kubeconfig"},
			Description: `Manage k3d kubeconfig`,
			Subcommands: []model.Subcommand{{
				Name: []string{"get", "print", "show"},
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to get kubeconfig for`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name: []string{"-a", "--all"},
				}},
			}, {
				Name:        []string{"merge", "write"},
				Description: `Write/Merge kubeconfig(s) from cluster(s) into new or existing kubeconfig/file`,
				Args: []model.Arg{{
					Name:        "cluster",
					Description: `Cluster to get kubeconfig for`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name: []string{"-a", "--all"},
				}, {
					Name:        []string{"-d", "--kubeconfig-merge-default"},
					Description: `Merge into the default kubeconfig ($KUBECONFIG or /Users/balli/.kube/config)`,
				}, {
					Name:        []string{"-s", "--kubeconfig-switch-context"},
					Description: `Switch to new context (default true)`,
				}, {
					Name:        []string{"-o", "--output"},
					Description: `Define output [ - | FILE ] (default from $KUBECONFIG or /Users/balli/.kube/config`,
					Args: []model.Arg{{
						Name:        "string",
						Description: `Define output [ - | FILE ] (default from $KUBECONFIG or /Users/balli/.kube/config`,
					}},
				}, {
					Name:        []string{"--overwrite"},
					Description: `[Careful!] Overwrite existing file, ignoring its contents`,
				}, {
					Name:        []string{"-u", "--update"},
					Description: `Update conflicting fields in existing kubeconfig (default true)`,
				}},
			}},
		}, {
			Name:        []string{"node"},
			Description: `Manage k3d nodes`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new node`,
				Args: []model.Arg{{
					Name:        "Node Name",
					Description: `Name of the node`,
				}},
				Options: []model.Option{{
					Name:        []string{"-c", "--cluster"},
					Description: `Cluster URL or k3d cluster name to connect to. (default "k3s-default")`,
					Args: []model.Arg{{
						Name:        "cluster",
						Description: `Cluster to connect to`,
						Generator:   nil, // TODO: port over generator
					}},
				}, {
					Name:        []string{"-i", "--image"},
					Description: `Node image to use`,
					Args: []model.Arg{{
						Name:        "image",
						Description: `Node image to use`,
					}},
				}, {
					Name:        []string{"--k3s-arg"},
					Description: `Additional k3s arguments`,
					Args: []model.Arg{{
						Name:        "k3s-arg",
						Description: `Additional k3s arguments`,
					}},
				}, {
					Name:        []string{"--k3s-node-label"},
					Description: `Specify k3s node labels in format`,
					Args: []model.Arg{{
						Name:        "k3s-node-label",
						Description: `Specify k3s node labels in format "foo=bar"`,
					}},
				}, {
					Name:        []string{"--memory"},
					Description: `Memory limit for the node container`,
					Args: []model.Arg{{
						Name:        "memory",
						Description: `Memory limit for the node container`,
					}},
				}, {
					Name:        []string{"-n", "--network"},
					Description: `Add node to (another) runtime network`,
					Args: []model.Arg{{
						Name: "network",
					}},
				}, {
					Name:        []string{"--replicas"},
					Description: `Number of replicas to create`,
					Args: []model.Arg{{
						Name:        "replicas",
						Description: `Number of replicas to create`,
						Suggestions: []model.Suggestion{{Name: []string{`1`}}, {Name: []string{`2`}}, {Name: []string{`3`}}, {Name: []string{`4`}}, {Name: []string{`5`}}},
					}},
				}, {
					Name:        []string{"--role"},
					Description: `Node role`,
					Args: []model.Arg{{
						Name:        "role",
						Description: `Node role`,
						Suggestions: []model.Suggestion{{Name: []string{`agent`}}, {Name: []string{`server`}}},
					}},
				}, {
					Name:        []string{"--runtime-label"},
					Description: `Specify runtime labels in format`,
					Args: []model.Arg{{
						Name:        "runtime-label",
						Description: `Specify runtime labels in format "foo=bar"`,
					}},
				}, {
					Name:        []string{"--timeout"},
					Description: `Maximum waiting time for '--wait' before canceling/returning`,
					Args: []model.Arg{{
						Name:        "timeout",
						Description: `Maximum waiting time for '--wait' before canceling/returning`,
					}},
				}, {
					Name:        []string{"-t", "--token"},
					Description: `Override cluster token`,
					Args: []model.Arg{{
						Name: "token",
					}},
				}, {
					Name:        []string{"--wait"},
					Description: `Wait for the node(s) to be ready before returning`,
				}},
			}, {
				Name:        []string{"delete"},
				Description: `Delete a node`,
				Args: []model.Arg{{
					Name:        "Node Name",
					Description: `Name of the node`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"-a", "--all"},
					Description: `Delete all existing nodes`,
				}, {
					Name:        []string{"-r", "--registry"},
					Description: `Also delete registries`,
				}},
			}, {
				Name:        []string{"edit", "update"},
				Description: `[EXPERIMENTAL] Edit node(s)`,
				Args: []model.Arg{{
					Name:        "Node Name",
					Description: `Name of the node`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name:        []string{"--port-add"},
					Description: `Map ports from the node container to the host`,
					Args: []model.Arg{{
						Name:        "port-add",
						Description: `Map ports from the node container to the host`,
					}},
				}},
			}, {
				Name:        []string{"list", "ls", "get"},
				Description: `List nodes`,
				Options: []model.Option{{
					Name:        []string{"--no-headers"},
					Description: `Don't print headers (default print headers)`,
				}, {
					Name:        []string{"-o", "--output"},
					Description: `Output format`,
					Args: []model.Arg{{
						Name:        "output",
						Description: `Output format`,
						Suggestions: []model.Suggestion{{Name: []string{`json`}}, {Name: []string{`yaml`}}},
					}},
				}},
			}, {
				Name:        []string{"start"},
				Description: `Start a node`,
				Args: []model.Arg{{
					Name:        "Node Name",
					Description: `Name of the node`,
					Generator:   nil, // TODO: port over generator
				}},
			}, {
				Name:        []string{"stop"},
				Description: `Stop a node`,
				Args: []model.Arg{{
					Name:        "Node Name",
					Description: `Name of the node`,
					Generator:   nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"registry", "registries", "reg"},
			Description: `Manage registry/registries`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"create"},
				Description: `Create a new registry`,
				Args: []model.Arg{{
					Name:        "Registry Name",
					Description: `Name of the registry`,
				}},
				Options: []model.Option{{
					Name:        []string{"--default-network"},
					Description: `Specify the network connected to the registry (default "bridge")`,
				}, {
					Name:        []string{"-i", "--image"},
					Description: `Specify image used for the registry (default "docker.io/library/registry:2")`,
					Args: []model.Arg{{
						Name: "image",
					}},
				}, {
					Name:        []string{"--no-help"},
					Description: `Disable the help text (How-To use the registry)`,
				}, {
					Name:        []string{"-p", "--port"},
					Description: `Select which port the registry should be listening on on your machine (localhost) (Format: [HOST:]HOSTPORT)`,
					Args: []model.Arg{{
						Name: "port",
					}},
				}, {
					Name:        []string{"--proxy-password"},
					Description: `Specify the password of the proxied remote registry`,
					Args: []model.Arg{{
						Name: "proxy-password",
					}},
				}, {
					Name:        []string{"--proxy-remote-url"},
					Description: `Specify the url of the proxied remote registry`,
					Args: []model.Arg{{
						Name: "proxy-remote-url",
					}},
				}, {
					Name:        []string{"--proxy-username"},
					Description: `Specify the username of the proxied remote registry`,
					Args: []model.Arg{{
						Name: "proxy-username",
					}},
				}, {
					Name:        []string{"-v", "--volume"},
					Description: `Mount volumes into the registry node (Format: [SOURCE:]DEST`,
					Args: []model.Arg{{
						Name: "volume",
					}},
				}},
			}, {
				Name:        []string{"delete", "del", "rm"},
				Description: `Delete registry/registries`,
				Args: []model.Arg{{
					Name:        "Registry Name",
					Description: `Name of the registry`,
					Generator:   nil, // TODO: port over generator
				}},
				Options: []model.Option{{
					Name: []string{"-a", "--all"},
				}},
			}, {
				Name:        []string{"list", "ls", "get"},
				Description: `List registries`,
				Args: []model.Arg{{
					Name:        "Registry Name",
					Description: `Name of the registry`,
					Generator:   nil, // TODO: port over generator
					IsOptional:  true,
				}},
				Options: []model.Option{{
					Name:        []string{"--no-headers"},
					Description: `Don't print headers (default print headers)`,
				}, {
					Name:        []string{"-o", "--output"},
					Description: `Output format`,
					Args: []model.Arg{{
						Name:        "output",
						Description: `Output format`,
						Suggestions: []model.Suggestion{{Name: []string{`json`}}, {Name: []string{`yaml`}}},
					}},
				}},
			}},
		}, {
			Name:        []string{"version"},
			Description: `Show k3d and default k3s version`,
		}},
	}
}
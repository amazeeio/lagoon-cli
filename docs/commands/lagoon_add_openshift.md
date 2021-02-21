## lagoon add openshift

Add a new openshift (platform admin user permissions only)

### Synopsis

Add a new openshift (platform admin user permissions only)

```
lagoon add openshift [flags]
```

### Options

```
  -C, --consoleUrl string         ConsoleURL of the openshift
  -h, --help                      help for openshift
  -I, --id int                    ID to assign the Openshift in Lagoon
  -m, --monitoringConfig string   Configuration for monitoring (JSON)
  -N, --name string               Name of the Openshift to be used in Lagoon
  -P, --projectUser string        Project user to use
  -R, --routerPattern string      Router pattern to use
  -S, --sshHost string            SSH host address
  -s, --sshPort string            SSH port number
  -T, --token string              Openshift token used to access the cluster
```

### Options inherited from parent commands

```
      --config-file string   Path to the config file to use (must be *.yml or *.yaml)
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force yes on prompts (if supported)
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
      --skip-update-check    Skip checking for updates
  -i, --ssh-key string       Specify path to a specific SSH key to use for lagoon authentication
```

### SEE ALSO

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments


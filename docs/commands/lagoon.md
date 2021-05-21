## lagoon

Command line integration for Lagoon

### Synopsis

Lagoon CLI. Manage your Lagoon hosted projects.

```
lagoon [flags]
```

### Options

```
      --config-file string   Path to the config file to use (must be *.yml or *.yaml)
      --debug                Enable debugging output (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force yes on prompts (if supported)
  -h, --help                 help for lagoon
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
      --skip-update-check    Skip checking for updates
  -i, --ssh-key string       Specify path to a specific SSH key to use for lagoon authentication
      --version              Version information
```

### SEE ALSO

* [lagoon add](lagoon_add.md)	 - Add a project, or add notifications and variables to projects or environments
* [lagoon config](lagoon_config.md)	 - Configure Lagoon CLI
* [lagoon delete](lagoon_delete.md)	 - Delete a project, or delete notifications and variables from projects or environments
* [lagoon deploy](lagoon_deploy.md)	 - Actions for deploying or promoting branches or environments in lagoon
* [lagoon export](lagoon_export.md)	 - Export lagoon output to yaml
* [lagoon fact](lagoon_fact.md)	 - Add and update facts
* [lagoon get](lagoon_get.md)	 - Get info on a resource
* [lagoon import](lagoon_import.md)	 - Import a config from a yaml file
* [lagoon kibana](lagoon_kibana.md)	 - Launch the kibana interface
* [lagoon list](lagoon_list.md)	 - List projects, deployments, variables or notifications
* [lagoon login](lagoon_login.md)	 - Log into a Lagoon instance
* [lagoon run](lagoon_run.md)	 - Run a task against an environment
* [lagoon ssh](lagoon_ssh.md)	 - Display the SSH command to access a specific environment in a project
* [lagoon update](lagoon_update.md)	 - Update a resource
* [lagoon version](lagoon_version.md)	 - Version information
* [lagoon web](lagoon_web.md)	 - Launch the web user interface
* [lagoon whoami](lagoon_whoami.md)	 - Whoami will return your user information for lagoon


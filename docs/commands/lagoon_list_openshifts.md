## lagoon list openshifts

List all Openshifts Lagoon knows about (platform admin user permissions only)

### Synopsis

List all Openshifts Lagoon knows about (platform admin user permissions only)
Only platform admin role can list openshifts.
NOTE: only openshifts that are used by projects can be listed

```
lagoon list openshifts [flags]
```

### Options

```
      --fields strings   Select which fields to display when listing Openshifts. Valid options (others are ignored): consoleurl,created,id,monitoringconfig,name,projectuser,routerpattern,sshhost,sshport,token (default [id,name])
  -h, --help             help for openshifts
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

* [lagoon list](lagoon_list.md)	 - List projects, deployments, variables or notifications


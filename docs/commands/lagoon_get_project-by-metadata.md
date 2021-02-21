## lagoon get project-by-metadata

Query lagoon projects by a given metadata key or key:value

### Synopsis

Query lagoon projects by a given metadata key or key:value

```
lagoon get project-by-metadata [flags]
```

### Options

```
  -h, --help           help for project-by-metadata
      --key string     The key name of the metadata value you are querying on
      --value string   The value for the key you are querying on
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

* [lagoon get](lagoon_get.md)	 - Get info on a resource


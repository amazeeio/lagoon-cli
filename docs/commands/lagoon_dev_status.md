## lagoon dev status

Report status of the pygmy services

### Synopsis

Loop through all of pygmy's services and identify the present state.
This includes the docker services, the resolver and SSH key status

```
lagoon dev status [flags]
```

### Options

```
  -h, --help   help for status
```

### Options inherited from parent commands

```
      --all-projects         All projects (if supported)
  -e, --environment string   Specify an environment to use
      --force                Force (if supported)
  -l, --lagoon string        The Lagoon instance to interact with
      --no-header            No header on table (if supported)
      --output-csv           Output as CSV (if supported)
      --output-json          Output as JSON (if supported)
      --pretty               Make JSON pretty (if supported)
  -p, --project string       Specify a project to use
  -i, --ssh-key string       Specify a specific SSH key to use
      --version              Version information
```

### SEE ALSO

* [lagoon dev](lagoon_dev.md)	 - start, stop or check the status of dev


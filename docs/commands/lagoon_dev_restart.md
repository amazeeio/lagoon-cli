## lagoon dev restart

Restart all pygmy containers.

### Synopsis

This command will trigger the Down and Up commands

```
lagoon dev restart [flags]
```

### Options

```
  -h, --help          help for restart
      --key string    Path of SSH key to add (default "/Users/ben/.ssh/id_rsa")
      --no-addkey     Skip adding the SSH key
      --no-resolver   Skip adding or removing the Resolver
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


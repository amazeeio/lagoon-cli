## lagoon fact add

Add a fact

### Synopsis

Add a fact

```
lagoon fact add [flags]
```

### Options

```
  -h, --help           help for add
  -N, --name string    The key name of the fact you are adding
  -V, --value string   The value of the fact you are adding
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

* [lagoon fact](lagoon_fact.md)	 - Add and update facts


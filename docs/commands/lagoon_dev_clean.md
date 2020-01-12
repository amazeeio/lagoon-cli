## lagoon dev clean

Stop and remove all pygmy services regardless of state

### Synopsis

Useful for debugging or system cleaning, this command will
remove all pygmy containers but leave the images in-tact.
This command does not check if the containers are running
because other checks do for speed convenience.

```
lagoon dev clean [flags]
```

### Options

```
  -h, --help   help for clean
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


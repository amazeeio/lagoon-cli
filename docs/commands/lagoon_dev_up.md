## lagoon dev up

Bring up pygmy services (dnsmasq, haproxy, mailhog, resolv, ssh-agent)

### Synopsis

Launch Pygmy - a set of containers and a resolver with very specific
configurations designed for use with Amazee.io local development.
It includes dnsmasq, haproxy, mailhog, resolv and ssh-agent.

```
lagoon dev up [flags]
```

### Options

```
  -h, --help          help for up
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


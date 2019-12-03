## lagoon update rocketchat

Update an existing rocketchat notification

### Synopsis

Update an existing rocketchat notification

```
lagoon update rocketchat [flags]
```

### Options

```
  -c, --channel string   The channel for the notification
  -h, --help             help for rocketchat
  -j, --json string      JSON string to patch
  -n, --name string      The current name of the notification
  -N, --newname string   The name of the notification
  -w, --webhook string   The webhook URL of the notification
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

* [lagoon update](lagoon_update.md)	 - Update project, environment, or notification

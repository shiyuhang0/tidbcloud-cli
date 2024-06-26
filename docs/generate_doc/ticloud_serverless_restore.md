## ticloud serverless restore

Restore a serverless cluster

```
ticloud serverless restore [flags]
```

### Examples

```
  Restore a serverless cluster in interactive mode:
 $ ticloud serverless restore

 Restore a serverless cluster with snaphot mode in non-interactive mode:
 $ ticloud serverless restore --backup-id <backup-id>

 Restore a serverless cluster with pointInTime mode in non-interactive mode:
 $ ticloud serverless restore -c <cluster-id> --backup-time <backup-time>
```

### Options

```
      --backup-id string     The ID of the backup. Used in snapshot restore mode.
      --backup-time string   The time to restore to (e.g. 2023-12-13T07:00:00Z). Used with point-in-time restore mode. Please specify the --cluster-id together.
  -c, --cluster-id string    The ID of cluster. Used in point-in-time restore mode. Please specify the --backup-time together.
  -h, --help                 help for restore
```

### Options inherited from parent commands

```
  -D, --debug            Enable debug mode
      --no-color         Disable color output
  -P, --profile string   Profile to use from your configuration file
```

### SEE ALSO

* [ticloud serverless](ticloud_serverless.md)	 - Manage TiDB Serverless clusters

###### Auto generated by spf13/cobra on 15-Apr-2024

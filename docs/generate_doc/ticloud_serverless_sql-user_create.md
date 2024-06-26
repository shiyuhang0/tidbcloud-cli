## ticloud serverless sql-user create

Create a SQL user

```
ticloud serverless sql-user create [flags]
```

### Examples

```
  Create a SQL user in interactive mode:
  $ ticloud serverless sql-user create

  Create a SQL user in non-interactive mode:
  $ ticloud serverless sql-user create --user <user-name> --password <password> --role <role> --cluster-id <cluster-id>
```

### Options

```
  -c, --cluster-id string   The ID of the cluster.
  -h, --help                help for create
      --password string     The password of the SQL user.
      --role strings        The role(s) of the SQL user.
  -u, --user string         The name of the SQL user.
```

### Options inherited from parent commands

```
  -D, --debug            Enable debug mode
      --no-color         Disable color output
  -P, --profile string   Profile to use from your configuration file
```

### SEE ALSO

* [ticloud serverless sql-user](ticloud_serverless_sql-user.md)	 - Manage cluster SQL users


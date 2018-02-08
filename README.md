# Backup and Restore SDK BOSH release

The Backup and Restore SDK BOSH release allows other BOSH deployed Cloud Foundry components to backup/restore their databases and blobstores.

**Docs**: [Release Author Guide](http://docs.cloudfoundry.org/bbr/bbr-devguide.html)

**Slack**: #bbr on https://slack.cloudfoundry.org

**Pivotal Tracker**: https://www.pivotaltracker.com/n/projects/1662777


## CI Status

Backup and Restore SDK Release status [![Build SDK Release Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/create-release/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release)

### Supported Databases

| Name     | Version | Status                                                                                                                                                                                                                                                                                     |
|:---------|:--------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| MariaDB  | 10.1.x  | [![MariaDB Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/mariadb-system-tests/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/mariadb-system-tests)            |
| MySQL    | 5.5.x   | [![MySQL Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/rds-mysql-5.5-system-tests/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/rds-mysql-5.5-system-tests)  |
| MySQL    | 5.6.x   | [![MySQL Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/rds-mysql-5.6-system-tests/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/rds-mysql-5.6-system-tests)  |
| MySQL    | 5.7.x   | [![MySQL Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/rds-mysql-5.7-system-tests/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/rds-mysql-5.7-system-tests)  |
| Postgres | 9.4.x   | [![Postgres Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/postgres-system-tests-9.4/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/postgres-system-tests-9.4) |
| Postgres | 9.6.x   | [![Postgres Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/postgres-system-tests-9.6/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/postgres-system-tests-9.6) |


### Supported Blobstores

| Name         | Status                                                                                                                                                                                                                                                                                                 |
|:-------------|:-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Versioned S3 | [![S3 Badge](https://backup-and-restore.ci.cf-app.com/api/v1/teams/main/pipelines/backup-and-restore-sdk-release/jobs/s3-blobstore-backuper-system-tests/badge)](https://backup-and-restore.ci.cf-app.com/teams/main/pipelines/backup-and-restore-sdk-release/jobs/s3-blobstore-backuper-system-tests) |

## Why?

Release authors wanting to write backup and restore scripts frequently need to back up and restore databases (or parts of databases).

Rather than have every team figure out the vagaries of backing up all the different kinds of database supported by CF, we've done it for you. The **Backup and Restore SDK** abstracts away the differences between databases, offering a consistent interface for your backup and restore scripts to use.

Behind the scenes, the SDK parses a configuration file passed to it, which selects the appropriate database backup/restore strategy (e.g. `pg_dump` or `mysql` at the required version) and places the backup artifact in the specified location.

## Config options

The SDK accepts a json document with the following fields

| name                  | type         | Optional | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
|:----------------------|:-------------|:---------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| username              | string       | no       | Database connection username                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| password              | string       | no       | Database connection password                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| host                  | string       | no       | Database connection host                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| port                  | integer      | no       | Database connection port, no defaulting is done, always needs to be specified                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| adapter               | string       | no       | Database adapter, see [Supported database adapters](#supported-database-adapters)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| database              | string       | no       | Name of the database to backup/restore                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| tables                | string array | yes      | If not specified, the entire database will be backed up/restored. If specified only the tables in that list will be included in the backup, and on restore the other tables in the database will be left as is. If the field is specified and empty, the utility will fail. If the field contains non-existent tables the utility will fail. We have not tested this with foreign key relationships or triggers spanning between tables specified in the `tables` list and other tables in the database not listed there. It's possible those relationships would be lost on restore. |
| tls.skip_host_verify  | bool         | yes      | Skip host verification for Server CA certificate                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| tls.certs.ca          | string       | yes      | Server CA certificate. This must be included if any of the `tls` block is specified                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| tls.certs.certificate | string       | yes      | Client certificate for Mutual TLS. This must be specified if `tls.certs.private_key` is given.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| tls.certs.private_key | string       | yes      | Client private key for Mutual TLS, this must be specified if `tls.certs.certificate` is given.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |

### Supported Database Adapters

* `postgres` (auto-detects versions between `9.4.x` and `9.6.x`)
* `mysql` (auto-detects `MariaDB 10.1.x`, and `MySQL 5.5.x`, `5.6.x`, `5.7.x`. Any other `mysql` variants are not tested)

## Deploying

### Deploying with `cf-deployment`

Users of [cf-deployment](https://github.com/cloudfoundry/cf-deployment) can simply apply the [backup-restore opsfile](https://github.com/cloudfoundry/cf-deployment/blob/master/operations/experimental/enable-backup-restore.yml). This will deploy the `database-backup-restorer` job on a backup restore VM alongside Cloud Foundry.

### Deploying as an instance group

You should co-locate the `database-backup-restorer` job and your release backup scripts on the same VM. If you use a dedicated backup-and-restore VM instance, co-locate them together on that VM. BOSH Lite is supported for testing.

Example BOSH v2 deployment manifest:
```yaml
...
instance_groups:
- name: backup
  networks:
  - name: my-network
  persistent_disk_type: 10GB
  stemcell: default
  update:
    serial: true
  vm_type: m3.large
  azs: [z1]
  instances: 1
  jobs:
  - name: backup-scripts
    properties:
      mydb:
      address: mydb.example.com
      db_scheme: mysql
      port: 3306
    release: my_release
  - name: database-backup-restorer
    release: backup-and-restore-sdk
...
```

## Usage from another BOSH job

### 1. Template `config.json`

Your job should template a `config.json` as follows:

```json
{
  "username": "db user",
  "password": "db password",
  "host": "db host",
  "port": 3306,
  "adapter": "db adapter; see 'Supported database adapters'",
  "database": "name of database to back up",
}
```
Or if you want to operate on specific tables

```json
{
  "username": "db user",
  "password": "db password",
  "host": "db host",
  "port": 3306,
  "adapter": "db adapter; see 'Supported database adapters'",
  "database": "name of database to back up",
  "tables": ["list", "of", "tables", "to", "back", "up"]
}
```

For the full list of `config.json` properties see [Config options](#config-options).

An example of templating using BOSH Links can be seen in the [cf networking release](https://github.com/cloudfoundry-incubator/cf-networking-release/blob/647f7a71b442c25ec29b1cc6484410946f41935c/jobs/bbr-cfnetworkingdb/templates/config.json.erb).


### 2. Write scripts to call the SDK binaries

In your release backup script, call `database-backup-restorer/bin/backup`:

```bash
/var/vcap/jobs/database-backup-restorer/bin/backup --config /path/to/config.json --artifact-file $BBR_ARTIFACT_DIRECTORY/artifactFile
```

In your release restore script, call `database-backup-restorer/bin/restore`:

```bash
/var/vcap/jobs/database-backup-restorer/bin/restore --config /path/to/config.json --artifact-file $BBR_ARTIFACT_DIRECTORY/artifactFile
```

The `restore` script will assume that the database schema has already been created, and matches the one of the backup. For BOSH releases, this usually means `restore` can be called after a successful deploy of the release, at the same version as the backup was taken.

#### Usage with [bbr](https://github.com/cloudfoundry-incubator/bosh-backup-and-restore)

For an example of the sdk being used in a release that can be backed up by bbr see the [exemplar release](https://github.com/cloudfoundry-incubator/exemplar-backup-and-restore-release).

### Developing
This repository using master as the main branch, tested releases are tagged with their versions.

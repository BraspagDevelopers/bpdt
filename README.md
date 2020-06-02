Braspag Deploy Tools
==========================
**Work in Progress**

Collection of tools to ease the deploy of applications


## `bpdt export-settings`
Convert multiples `appsettings.*.json` files to `.env` file syntax

### Usage
```
bpdt export-settings [-d <dir>] -f <file1> -f <file2>
```

#### Flags
- **`--directory`, `-d`:** Directory where the files will be looked for
- **`--file`, `-f`:** Files that will be used as input

## `bpdt patch-nuget`
Add clear text passwords to a nuget config file

### Usage
```
bpdt export-nuget <path> <nugetSource> <username> <password>
```

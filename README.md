`bpdt` - Braspag Deploy Tools
==========================
**Work in Progress**

Collection of tools to ease the deploy of applications



## Installation

### [gobinaries.com](gobinaries.com) method

Install to `/usr/local/bin`
```bash
curl -sf https://gobinaries.com/BraspagDevelopers/bphc | sh
```

You can also specify a custom directory where to download the binary file
```bash
# Install on the current directory
curl -sf https://gobinaries.com/BraspagDevelopers/bpdt | PREFIX=. sh
```
```bash
# Install on /tmp
curl -sf https://gobinaries.com/BraspagDevelopers/bpdt | PREFIX=/tmp sh
```

### `go get` method
```bash
go get github.com/BraspagDevelopers/bpdt
```

### Usage
#### `bpdt export-settings`
Convert multiples `appsettings.*.json` files to `.env` file syntax

```
bpdt export-settings [-d <dir>] -f <file1> -f <file2>
```

###### Flags
- **`--directory`, `-d`:** Directory where the files will be looked for
- **`--file`, `-f`:** Files that will be used as input

##### `bpdt patch-nuget`
Add clear text passwords to a nuget config file

```
bpdt export-nuget <path> <nugetSource> <username> <password>
```

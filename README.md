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

## Usage
### `bpdt export-settings`
Convert multiples `appsettings.*.json` files to `.env` file syntax

```bash
bpdt export-settings [-d <dir>] -f <file1> -f <file2>
```

#### Flags
- **`--directory`, `-d`:** Directory where the files will be looked for
- **`--file`, `-f`:** Files that will be used as input

### `bpdt patch-nuget`
Add clear text passwords to a nuget config file

```bash
bpdt export-nuget <path> <nugetSource> <username> <password>
```

### `bpdt env-to-yaml`
Add entries to a YAML element using a .env file as input

```bash
bpdt env-to-yaml <.env-file-path> <yaml-file-path>
```

#### Arguments
1. **`<.env-file-path>`:** Path of the _.env_ file
2. **`<yaml-file-path>`:** Path of the YAML file

#### Flags
- **`--directory`, `-d`:** Directory where the files will be looked for
- **`--ypath`:** A period separated string indicating where in the YAML the variables should be appended


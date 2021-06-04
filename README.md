![License](https://img.shields.io/github/license/BraspagDevelopers/bpdt.svg?style=flat-square)
![Tag](https://img.shields.io/github/tag/BraspagDevelopers/bpdt.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/BraspagDevelopers/bpdt?style-flat-square)](https://goreportcard.com/report/github.com/BraspagDevelopers/bpdt)


`bpdt` - Braspag Deploy Tools
==========================
**Work in Progress**

Collection of tools to ease the deploy of applications



## Installation

### [gobinaries.com](https://gobinaries.com) method

Install to `/usr/local/bin`
```bash
curl -sf https://gobinaries.com/BraspagDevelopers/bpdt | sh
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
go get -u github.com/BraspagDevelopers/bpdt
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
bpdt patch-nuget <path> <nugetSource> <username> <password>
```
#### Arguments
1. **`<path>`:** Path of the nuget config file
2. **`<nugetSource>`:** Nuget source to which add the credentials
3. **`<username>`:** Username
4. **`<password>`:** Password

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


### `bpdt ref-secrets`
Adds a secret key reference variable to a yaml file.
The file is tipically a kubernetes Deployment or Pod resource file.

```bash
bpdt ref-secrets <file-path> <secret-name>
```

#### Arguments
1. **`<file-path>`:** Path of the _.env_ file
2. **`<secret-name>`:** Path of the YAML file

#### Flags
- **`--ypath`:** A period separated string indicating where in the YAML the variables are placed. _Default: `spec.template.spec.containers.0.env`_
- **`--directory`, `-d`:**: Directory where the files will be looked for. _Default: working directory_

- **`--prefix`, `-p`:** The prefix for the secret variables. _Default: `#<Secret>{`_
- **`--suffix`, `-s`:** The suffix for the secret variables. _Default: `}#`_


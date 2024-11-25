# Create Certificate Script

Creates the following files
* `<target_directory>/<name>_key`
* `<target_directory>/<name>_cert`
* `<target_directory>/<name>_ski`

This files can be used to start an eebus instance.

## Usage

```sh
go run create_cert/main.go <target_directory> <name>
```

* `target_directory`: The path to the directory to put the files into
* `name`: Prefix of the generated files.

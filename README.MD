# smg

    A cli tool, that helps me automate shit on server, at least that's the idea.

## usage 

```bash

$bash /home/ubuntu > smg


Usage:
  smg [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  dbDump      dbDump filename.sql
  decrypt     Decrypts string/file
  encrypt     Encrypts string/file
  gen         Generate the hash from string/file
  gen         Generate the hash from string/file
  genkeys     Generates RSA key pair
  help        Help about any command
  lump        lump nginx mysql php
  verify      Verify the hash from string/file

Flags:
  -h, --help   help for smg

Use "smg [command] --help" for more information about a command. 

```


# Available Commands

 - lump
 - dbDump 
 - decrypt
 - encrypt
 - gen
 - verify

## lump

> lump allows you to install a few req. tools without looking into docs.

> ###  Supported Application and their identifiers
> - Php ( php7.4 , php7.3 , php7.2 ) 
> - Mysql ( mysql ) 
> - Nginx ( nginx )

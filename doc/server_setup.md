# Setting up projectxserver

## The easy way

XX

## Conventions
Throughout this document, we will mark commands to be run on your
local machine with the shell prompt `local$` and commands to be
run on your server with `server%`.

For example:

```
local$ projectx signup -server=projectx.example.com you@gmail.com
```
and
```
server% sudo systemctl stop projectxserver.service
```

## Introduction
This document describes the process for creating an ProjectX installation by deploying
an `projectxserver`, a combined ProjectX Store and Directory server, to
a Linux-based machine.

The installation will use the central ProjectX key server (`key.projectx.io`) for
authentication, which permits inter-operation with other ProjectX servers.

There are multiple versions of `projectxserver`, each depending on where the
associated storage is kept, either on the server's local disk or with a cloud
storage provider.
The binaries that use cloud storage providers each have a suffix that
identifies the provider, such as `projectxserver-gcp` for the Google Cloud
Platform.
These binaries are also kept in distinct repositories, such as `gcp.projectx.io`
for the Google Cloud Platform.

The process follows these steps:

- [sign up](#signup) for an ProjectX user account
- [configure](#domain) a domain name and create an ProjectX user for the server,
- if necessary, [set up the cloud](#cloud
) storage service,
- [deploy](#deploy) the `projectxserver` to a Linux-based server,
- [configure](#configure) the `projectxserver`.

Each of these steps (besides deployment) has a corresponding `projectx`
subcommand to assist you with the process.

## Prerequisites

To deploy an `projectxserver` you need to decide on values for:

- An Internet domain to which you can add DNS records.
  (We will use `example.com` in this document.)
  Note that the domain need not be dedicated to your ProjectX installation; it
  just acts as a name space inside which you can create ProjectX users for
  administrative purposes.

- Your ProjectX user name (an email address).
  (We will use `you@gmail.com` in this document.)
  This user will be the administrator of your ProjectX installation.
  The address may be under any domain,
  as long you can receive mail at that address.

- The host name of the server on which `projectxserver` will run.
  (We will use `projectx.example.com` in this document.)

## XX {#XX}

To register your public key with the central key server run `projectx signup`,
passing your chosen host name as its `-server` argument
and your chosen ProjectX user name as its final argument.
Then follow the onscreen instructions.

The [Signing up a new user](/doc/signup.md) document describes this process in
detail.
If you change your mind about the host name, you can update with `projectx user -put`.

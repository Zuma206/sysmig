# ⏩ sysmig

> [!WARNING]
> This project is in an alpha testing stage. Feel free to try the project out, but issues and breaking changes may occur.

> [!NOTE]
> Full documentation is coming soon.

Declarative system configuration for POSIX operating systems

## Installation

### Install script

It is currently recommended to install sysmig using the official install script. In future, sysmig may also be able to self-update, making this a full solution.

#### Using curl

```sh
sudo curl -o /usr/local/bin/sysmig https://github.com/Zuma206/sysmig/releases/download/v0.0.0-fix-action-permissions/sysmig
sudo chmod +x /usr/local/bin/sysmig
```

### Package Managers

By design, sysmig is not and will never be officially packaged for any package manager. This is by design, as sysmig is expected to install and/or configure your package managers in a migration, and therefore should be installed outside of them.

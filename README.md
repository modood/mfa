mfa
===

[![license](https://img.shields.io/badge/license-WTFPL%20--%20Do%20What%20the%20Fuck%20You%20Want%20to%20Public%20License-green.svg)](https://github.com/modood/mfa/blob/master/LICENSE)

Generate TOTP(Time-based One-time Password) token on the command line.

Configuration
-------------

Default config file is $HOME/.mfa.json

```json
{
  "google": "your secret key",
  "github": "your secret key"
}
```

Installation
------------

```
$ go install github.com/modood/mfa@latest
```

License
-------

This repo is released under the [WTFPL](http://www.wtfpl.net/) – Do What the Fuck You Want to Public License.

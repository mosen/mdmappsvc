# mdmappsvc
A micro service for storing information about catalogs of installable applications and media on disparate sources.
It should also generate install manifests suitable for the MDM InstallApplication command targeted at macOS.


## Notes ##

### Sources ###

Sources represent individual package and media sources. They are uniquely defined by the protocol used for r/w access to
the source, and the set of credentials needed to scan the source.

One example might be a Munki repository, which may be served to clients over HTTPS, but which requires a set of credentials
and an alternate access method to enumerate the contents of the repository.

### Checksums ###

Not all sources can provide their own checksum information about package components. The OSX mdm client requires
chunked checksums however. This service shall be responsible for generating MD5 checksums for chunked transfers.


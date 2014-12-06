purl
====

Perl, but fluffy like a :cat:

In Go of course!

- Embedded Perl 5.20.1
- Call native Go functions via XS
- Evaluates any Perl code with persistent state

Thanks to [Campher](https://github.com/bradfitz/campher) for a pointer in the right direction :smirk:

### Building purl

Currently needs an absolute path workaround:

- Edit [perl/perl.go](perl/perl.go) and update -L path to match yours
- Run `make`
- Run `./purl`

### Licence

Copyright ©‎ 2014, Ian Kent (http://iankent.uk).

Released under MIT license, see [LICENSE](LICENSE.md) for details.

# FHIR Core (Go)

![Continuous Integration](https://img.shields.io/github/actions/workflow/status/friendly-fhir/go-fhir/.github%2Fworkflows%2Fpostsubmit.yaml?logo=github)
[![GitHub Release](https://img.shields.io/github/v/release/friendly-fhir/go-fhir?include_prereleases)][github-releases]
[![Gitter Channel](https://img.shields.io/badge/matrix-%23friendly--fhir-darkcyan?logo=gitter)][gitter-channel]
[![readthedocs](https://img.shields.io/badge/docs-readthedocs-blue?logo=readthedocs&logoColor=white)][docs]
[![Godocs](https://img.shields.io/badge/docs-godocs-blue?logo=go&logoColor=white)][go-docs]

> [!NOTE]
> The API for these generated bindings are current **unstable**.
> These have been published primarily so that the [`fhenix`] utility may
> leverage them to bootstrap its own parsing of `StructureDefinition`s.

This repository contains (mostly) generated [Go] bindings for the FHIR core
specification. This can be leveraged by other projects to parse and model
FHIR resources in a Go-idiomatic way.

Code bindings have been generated using the [`fhenix`] utility.

[Go]: https://golang.org/

[gitter-channel]: https://matrix.to/#/#friendly-fhir:gitter.im
[docs]: https://friendly-fhir.github.io/go-fhir/
[go-docs]: https://pkg.go.dev/github.com/friendly-fhir/go-fhir
[github-releases]: https://github.com/friendly-fhir/go-fhir/releases
[`fhenix`]: https://github.com/friendly-fhir/fhenix

## Installation

```bash
go get github.com/friendly-fhir/go-fhir@latest
```

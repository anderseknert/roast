# Changelog

## [0.4.1] - 2024-10-01

### Changed

- Update actual dependencies used (i.e. `go mod tidy`)

### Changed

## [0.4.0] - 2024-10-01

### Changed

- New location format 
- Removed `name` attribute from rules in favor of using the rule's `ref` to infer name
- Updated OPA version from v0.68.0 to v0.69.0

## [0.3.0] - 2024-09-25

### Changed

- Removed `annotations` from module, in favor of annotations on `package` and `rules`

## [0.2.0] - 2024-09-09

### Changed

- OPA version updated from v0.67.1 to v0.68.0

## [0.1.1] - 2024-09-09

### Changed

- Fixed issue in annotations encoding, where multiple `custom` attributes wouldn't be encoded
  with a `,` separator.

## [0.1.0] - 2024-08-20

First release!

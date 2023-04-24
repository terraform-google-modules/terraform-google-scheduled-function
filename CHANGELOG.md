# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][keepachangelog-site],
and this project adheres to [Semantic Versioning][semver-site].

## [2.5.1](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.5.0...v2.5.1) (2023-04-06)


### Bug Fixes

* **deps:** update module golang.org/x/net to v0.7.0 [security] ([#98](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/98)) ([8fd376d](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/8fd376d477444e36951d407c62d60c2d98f1c337))
* fixes for tflint ([#110](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/110)) ([098214d](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/098214d3e3b5a6220f46589f9467e607c8a04b67))
* update function_max_instances to null default ([#106](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/106)) ([8f32217](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/8f32217ba0570ef2b1476b1e1f6ae1728c2cae12))

## [2.5.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.4.0...v2.5.0) (2022-06-09)


### Features

* removal of folder should use the same time constraint used for projects ([#79](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/79)) ([f4ab92b](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/f4ab92b878ef75b961ab971e5ef31f82d76ee768))
* support deletion of firewall policies attached to folders in clean up module ([#76](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/76)) ([87037dd](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/87037ddbb93534ed173a9cc902eb674f7056333e))

## [2.4.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.3.0...v2.4.0) (2022-05-09)


### Features

* add topic_labels ([#73](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/73)) ([e00736a](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/e00736a2ac2aa8a5794d86ce032142db037a7a2e))

## [2.3.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.2.0...v2.3.0) (2022-02-15)


### Features

* add function max instances argument ([#66](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/66)) ([#67](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/67)) ([f08a0a4](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/f08a0a4552fd192e1ee88c59f04d1a343ac89910))

## [2.2.0](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.1.0...v2.2.0) (2021-11-16)


### Features

* update TPG version constraints to allow 4.0 ([#63](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/63)) ([248b914](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/248b914e7d7c4a384224c862d98b0846b0cdda7c))

## [2.1.0](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.0.0...v2.1.0) (2021-08-13)


### Features

* Add option for VPC connector ([#59](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/59)) ([9a82b83](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/9a82b83f42602123dcf88a8e518c7aaa6deac707))

## [2.0.0](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.6.0...v2.0.0) (2021-04-29)


### ⚠ BREAKING CHANGES

* add Terraform 0.13 constraint and module attribution (#50)

### Features

* add Terraform 0.13 constraint and module attribution ([#50](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/50)) ([3bb8cba](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/3bb8cba3170252d4b390525dfcbfab20cc9b4531))


### Bug Fixes

* Add folder editor permissions to delete folders for cleanup submodule ([#55](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/55)) ([f9e3841](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/f9e3841556f35a621d8d50530f9591d88f090dd8))

## [1.6.0](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.5.1...v1.6.0) (2021-04-05)


### Features

* add folder clean up ([#47](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/47)) ([6d001bb](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/6d001bb24197e6475500cdcfc8c291aabe41699c))
* add retry logic for quota errors ([#53](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/53)) ([293ed2c](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/293ed2c7f77ba14e0fd6d1d2cb01b08d9aa58968))


### Bug Fixes

* add grant_token_creator flag for pubsub  ([#52](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/52)) ([7fee659](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/7fee659322dd014b818a26cc0c132b2b71ca91d2))

### [1.5.1](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.5.0...v1.5.1) (2020-09-02)


### Bug Fixes

* default bucket name behavior ([#42](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/42)) ([5adad84](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/5adad84a4af7f58f6934779500f9bf3fa38464a1))

## [1.5.0](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.2...v1.5.0) (2020-08-11)


### Features

* Update project cleaner function to clean up Cloud Endpoints ([#38](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/38)) ([4403dbe](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/4403dbe45e27c86f34928550ad44beff1e92a92b))

### [1.4.2](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.1...v1.4.2) (2020-05-06)


### Bug Fixes

* Add Terraform 0.12 support to bucket_force_destroy boolean variable ([#24](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/issues/24)) ([2f11dfa](https://www.github.com/terraform-google-modules/terraform-google-scheduled-function/commit/2f11dfab523d82fa925ddbdf69fe6df8299f98d6))

## [Unreleased]

## [1.4.1] - 2020-01-13

### Fixed

- Fixed issue with scheduler job output value inconsistency. [#29]

## [1.4.0] - 2019-12-19

### Added

- The `scheduler_job` variable, the `scheduler_job` output, and the `pubsub_topic_name` output to enable linking multiple
  instances of the module to the same Cloud Scheduler job. [#15]

## [1.3.0] - 2019-12-18

### Added

- The `function_source_dependent_files` variable is passed on to the `event-function` module's `source_dependent_files` variable. [#28]

### Changed

- The function implementation is provided by the Event Function module. [#6]

## [1.2.0] - 2019-11-20

### Added

- The `function_timeout_s` variable is exposed on the `project_cleanup` submodule.

## [1.1.1] - 2019-11-13

### Fixed

- The IAM module was replaced with IAM member resources to support dynamic members in additive mode. [#22]

## [1.1.0] - 2019-11-11

### Changed

- The `project_cleanup` submodule can be scheduled to remove labelled or unlabelled projects. [#20] [#21]

### Added

- The `logs-slack-alerts` example. [#13]

## [1.0.0] - 2019-07-30

### Changed

- Supported version of Terraform is 0.12. [#11]

## [0.4.1] - 2019-07-03

### Fixed

- Project and region are applied to the scheduler job. [#8]

## [0.4.0] - 2019-06-17

### Added

- A variable which configures the time zone of the scheduler job. [#5]

## [0.3.0] - 2019-06-11

### Added

- Submodule which cleans up old projects.

## [0.2.0] - 2019-04-02

### Added

- Ability to specify a service account for functions to run as

## [0.1.0] - 2019-03-14

### Added

- Initial release

[Unreleased]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.1...HEAD
[1.4.1]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.4.1...v1.0.0
[0.4.1]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/releases/tag/v0.1.0

[#29]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/29
[#28]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/28
[#22]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/22
[#21]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/21
[#20]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/20
[#15]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/15
[#13]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/13
[#11]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/11
[#8]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/8
[#6]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/6
[#5]: https://github.com/terraform-google-modules/terraform-google-scheduled-function/pull/5

[keepachangelog-site]: https://keepachangelog.com/en/1.0.0/
[semver-site]: https://semver.org/spec/v2.0.0.html

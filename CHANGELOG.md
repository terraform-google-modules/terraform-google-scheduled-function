# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][keepachangelog-site],
and this project adheres to [Semantic Versioning][semver-site].

## [4.2.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v4.1.0...v4.2.0) (2024-05-22)


### Features

* Support deletion of Billing Account Log Sinks ([#227](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/227)) ([e58dc20](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/e58dc20d35745996b9496709e5ed0de43bc558a3))


### Bug Fixes

* **deps:** Update go modules ([#221](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/221)) ([f962d8e](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/f962d8e8e4cbcac7fc46f156932e0d24583a6e78))
* **deps:** Update go modules ([#230](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/230)) ([d9d73b0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/d9d73b0e8b4829527210761ea235a5d81a13cc07))

## [4.1.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v4.0.1...v4.1.0) (2024-05-10)


### Features

* support deletion of Cloud Asset Inventory feeds not in use in organization ([#198](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/198)) ([bd8f8ad](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/bd8f8ad694edecb290ef8ae39ac5900477c8f6bb))
* support deletion of SCC Notification not in use in organization ([#196](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/196)) ([b0e3ba0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/b0e3ba0d3f8e115bfb16bda53352b9acd9c3113a))


### Bug Fixes

* **deps:** Update go modules ([#217](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/217)) ([ac847ad](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/ac847ad22102beafdb710ca0eacae899e17271fd))
* **deps:** Update module cloud.google.com/go/asset to v1.19.0 ([#215](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/215)) ([8aecc2d](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/8aecc2df0c1d731281e7574f9d1fec1ce08aad17))
* **deps:** Update module cloud.google.com/go/securitycenter to v1.29.0 ([#212](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/212)) ([0fdb546](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/0fdb5464e14c84bbd10523b0dabbe0a8cfbec2f4))
* **deps:** Update module golang.org/x/oauth2 to v0.20.0 ([#213](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/213)) ([6695683](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/66956831e42ae452288ff908c7bf7a6332a7bd50))

## [4.0.1](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v4.0.0...v4.0.1) (2024-05-01)


### Bug Fixes

* delete even single endpoint service ([#207](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/207)) ([423ec36](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/423ec36e0fb7ddb046e3ef6338bbfb18387cbacf))
* **deps:** Update go modules ([#185](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/185)) ([7da1975](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/7da1975b53c4b004a02db2f128e71d656036d570))
* **deps:** Update go modules ([#195](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/195)) ([03ad588](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/03ad58860b0402efd888e79c3286facaebd8ddfd))
* **deps:** Update go modules ([#205](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/205)) ([b8de15f](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/b8de15f73b66648d7ab6585263e6fca5543c89e4))
* **deps:** Update go modules and dev-tools ([#199](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/199)) ([b94bede](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/b94bede8f6da64509dfd36c5c1e378882a37d414))
* **deps:** Update module golang.org/x/net to v0.24.0 ([#197](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/197)) ([35e4084](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/35e4084f97350f3c4101b5e920b666554688bd52))
* updates for go modules ([#184](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/184)) ([ff0be2f](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/ff0be2f2d2b8af2de01ea2601c5aa91edd1c0b2e))

## [4.0.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v3.0.0...v4.0.0) (2024-01-30)


### ⚠ BREAKING CHANGES

* **GO1.21:** support deletion of Tag Keys not in use in organizations in the clean up module ([#175](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/175))
* **TPG >=4.23:** Add function_docker_registry variables and use them in terraform-google-event-function module ([#150](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/150))

### Features

* add create_bucket flag ([#147](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/147)) ([d2b68fc](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/d2b68fc86c433a0e42a524bf7f5ae8defc2d6522))
* **GO1.21:** support deletion of Tag Keys not in use in organizations in the clean up module ([#175](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/175)) ([6cc42cd](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/6cc42cd3ef23ebd1b17c6cc4ec259ea78ed30d82))
* **TPG >=4.23:** Add function_docker_registry variables and use them in terraform-google-event-function module ([#150](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/150)) ([5c9ddcd](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/5c9ddcdc707543d2ad144d00dbbb1f0a33bf25d6))


### Bug Fixes

* **deps:** Bump golang.org/x/crypto from 0.16.0 to 0.17.0 in /modules/project_cleanup/function_source ([#173](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/173)) ([1d12e4f](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/1d12e4f91f06dd0c352c3fd8e97f5880ca07ec0c))
* **deps:** Update GO modules ([#165](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/165)) ([ccb1961](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/ccb1961b978a9ca51858bdbc901faa15ae4b0b62))
* **deps:** Update GO modules ([#166](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/166)) ([4796116](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/4796116962c2a14336e0da6968f45f3f7abc2cab))
* **deps:** update go modules ([#176](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/176)) ([77c764a](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/77c764ae580117ec561f784945b5d1eaa0a041a4))
* **deps:** Update module google.golang.org/api to v0.149.0 ([#163](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/163)) ([5ba8c74](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/5ba8c744975069d06e20722f2d1a0d6804a967a5))

## [3.0.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.6.0...v3.0.0) (2023-11-01)


### ⚠ BREAKING CHANGES

* updates pubsub module to v6 ([#159](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/159))
* **deps:** Update GO modules ([#126](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/126))
* **deps:** update go to 1.20 ([#132](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/132))
* **deps:** Update TF modules (major) ([#142](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/142))

### Features

* add topic kms key name variable ([#145](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/145)) ([667c4b8](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/667c4b8e45f34b85a42fb3b3f79683f782117183))


### Bug Fixes

* Add required api to README ([#139](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/139)) ([d046d2b](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/d046d2bcc224aa0ed26fd14b954eeb10c23ddcc8))
* **deps:** Update GO modules ([#126](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/126)) ([c544227](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/c544227ad14a7f2004b7689d310e8fb24ee25092))
* **deps:** update go to 1.20 ([#132](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/132)) ([1f95f6b](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/1f95f6bcb19ad99c92f43134923b488ef36dd878))
* **deps:** Update TF modules (major) ([#142](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/142)) ([f1cf69e](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/f1cf69eeb1cf21fd576b96beb9a5c05a3e736834))
* updates pubsub module to v6 ([#159](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/159)) ([0b3ea07](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/0b3ea075014b43dd4c3470439559fc5575599db2))

## [2.6.0](https://github.com/terraform-google-modules/terraform-google-scheduled-function/compare/v2.5.1...v2.6.0) (2023-06-15)


### Features

* include Ingress Settings variable ([#129](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/129)) ([a68339d](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/a68339d325ff5dd16ca7f97fcdeee100610a8191))
* support adding secrets from secret manager ([#123](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/123)) ([5b8f226](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/5b8f2267016b613234b20a7174ce54242fe5e23f))


### Bug Fixes

* **deps:** update go modules ([#121](https://github.com/terraform-google-modules/terraform-google-scheduled-function/issues/121)) ([1acd373](https://github.com/terraform-google-modules/terraform-google-scheduled-function/commit/1acd37393c860889945ded66deb25d3a387986b0))

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

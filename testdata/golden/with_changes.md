## Summary of Fleet Package Changes

Report generated from snapshot branch commit
[fe2bc77a4f4294b236899b9b9c61d2d118e059b0](
https://github.com/elastic/package-storage/commit/fe2bc77a4f4294b236899b9b9c61d2d118e059b0)
from 2022-01-17 11:09:03 &#43;0000 UTC.

Comparisons were made to production branch commit
[abc1234](
https://github.com/elastic/package-storage/commit/abc1234)
from 2022-01-11 08:09:18 &#43;0000 UTC.

Filtering parameters:

  - Include Deprecated: false


### AWS - 1.8.0
Owner: elastic/foo-team

Requires: ^8.0.0

Changes since 1.7.0

  - 1.8.0
     - enhancement: Update ECS ([PR](https://github.com/elastic/integrations/pull/123))
     - bugfix: Fix bug ([PR](https://github.com/elastic/integrations/pull/124))
  




To promote these packages use this command:

`elastic-package promote -d=snapshot-production -n -p "aws-1.8.0"`

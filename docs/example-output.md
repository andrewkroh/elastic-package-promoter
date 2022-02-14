## Summary of Fleet Package Changes

Report generated from snapshot branch commit
[1a8a5f8bfd2a91c848587b04553eb9df09948bde](
https://github.com/elastic/package-storage/commit/1a8a5f8bfd2a91c848587b04553eb9df09948bde)
from 2022-02-14 12:41:42 &#43;0000 UTC.

Comparisons were made to production branch commit
[4667bfc5e46a7e16cc1bf5a5afda81e1b9a1e1d1](
https://github.com/elastic/package-storage/commit/4667bfc5e46a7e16cc1bf5a5afda81e1b9a1e1d1)
from 2022-02-11 19:36:14 &#43;0000 UTC.

Filtering parameters:

  - Team: elastic/integrations

  - Include Deprecated: false


### Apache HTTP Server - 1.3.4
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.3.2

  - 1.3.4
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 1.3.3
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### AWS - 1.11.4
Owner: elastic/integrations

Requires: ^7.15.0 || ^8.0.0

Changes since 1.11.0

  - 1.11.4
     - bugfix: Add Ingest Pipeline script to map IANA Protocol Numbers ([PR](https://github.com/elastic/integrations/pull/2470))
  
  - 1.11.3
     - bugfix: Changing missing ecs versions to 8.0.0 ([PR](https://github.com/elastic/integrations/pull/2642))
  
  - 1.11.2
     - bugfix: Add data_stream.dataset option for custom aws-cloudwatch log input ([PR](https://github.com/elastic/integrations/pull/2560))
  
  - 1.11.1
     - bugfix: Update permission list ([PR](https://github.com/elastic/integrations/pull/2635))
  



### Custom AWS Logs - 0.1.0
Owner: elastic/integrations

Requires: ^7.16.0 || ^8.0.0

New Package

  - 0.1.0
     - enhancement: initial release ([PR](https://github.com/elastic/integrations/pull/2353))
  

### Azure Resource Metrics - 1.0.2
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.0.1

  - 1.0.2
     - enhancement: Update documentation ([PR](https://github.com/elastic/integrations/pull/2656))
  



### Cassandra - 1.2.0
Owner: elastic/integrations

Requires: ^7.15.0 || ^8.0.0

Changes since 1.1.0

  - 1.2.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2483))
  



### CockroachDB Metrics - 0.2.1
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 0.2.0

  - 0.2.1
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2484))
  



### etcd - 0.1.1
Owner: elastic/integrations

Requires: ^7.15.0 || ^8.0.0

New Package

  - 0.1.1
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2486))
  
  - 0.1.0
     - enhancement: Add metrics dashboard ([PR](https://github.com/elastic/integrations/pull/2336))
  
  - 0.0.1
     - enhancement: Initial release ([PR](https://github.com/elastic/integrations/pull/2167))
  

### HAProxy - 1.1.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 0.7.0

  - 1.1.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2488))
  
  - 1.0.2
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 1.0.1
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  
  - 1.0.0
     - enhancement: Release HAProxy as GA ([PR](https://github.com/elastic/integrations/pull/1608))
  



### IIS - 0.8.2
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 0.8.0

  - 0.8.2
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 0.8.1
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### Kafka - 1.2.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.1.0

  - 1.2.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2490))
  



### Kubernetes - 1.17.0
Owner: elastic/integrations

Requires: ^7.16.0 || ^8.0.0

Changes since 1.9.0

  - 1.17.0
     - enhancement: Disable audit logs collection by default ([PR](https://github.com/elastic/integrations/pull/))
  
  - 1.16.0
     - enhancement: Documentation improvements ([PR](https://github.com/elastic/integrations/pull/2657))
  
  - 1.15.0
     - enhancement: Add ssl.certificate_authorities configuration ([PR](https://github.com/elastic/integrations/pull/2613))
  
  - 1.14.3
     - bugfix: Add missing job.name and cronjob.name fields to state_container datastream ([PR](https://github.com/elastic/integrations/pull/2625))
  
  - 1.14.2
     - bugfix: Add missing job.name and cronjob.name fields to container related datastreams ([PR](https://github.com/elastic/integrations/pull/2612))
  
  - 1.14.1
     - bugfix: Add missing job.name and cronjob.name fields added by metadata generators ([PR](https://github.com/elastic/integrations/pull/2608))
  
  - 1.14.0
     - enhancement: Tune state_metrics settings ([PR](https://github.com/elastic/integrations/pull/2567))
  
  - 1.13.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2491))
  
  - 1.12.0
     - enhancement: Expose add_recourse_metadata configuration option ([PR](https://github.com/elastic/integrations/pull/2370))
  
  - 1.11.0
     - enhancement: Add memory.working_set.limit.pct for pod and container data streams ([PR](https://github.com/elastic/integrations/pull/2469))
  
  - 1.10.0
     - enhancement: Add leader election in state_job data stream ([PR](https://github.com/elastic/integrations/pull/2377))
  



### MySQL - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.1

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2496))
  



### NATS - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2504))
  
  - 1.2.1
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### Nginx - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.1

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2505))
  
  - 1.2.3
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 1.2.2
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### Nginx Ingress Controller Logs - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2506))
  
  - 1.2.2
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 1.2.1
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### Prometheus Metrics - 0.9.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 0.7.0

  - 0.9.0
     - enhancement: Add standard HTTP options to the package ([PR](https://github.com/elastic/integrations/pull/2632))
  
  - 0.8.0
     - enhancement: Improve default datastream enablement ([PR](https://github.com/elastic/integrations/pull/2619))
  



### RabbitMQ Logs - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2509))
  



### Redis - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2510))
  



### STAN - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2511))
  



### System - 1.10.0
Owner: elastic/integrations

Requires: ^7.16.0 || ^8.0.0

Changes since 1.6.4

  - 1.10.0
     - enhancement: Expose winlog input ignore_older option. ([PR](https://github.com/elastic/integrations/pull/2542))
     - bugfix: Fix preserve original event option ([PR](https://github.com/elastic/integrations/pull/2542))
     - enhancement: Make order of Security, Application, System options consistent with other winlog based integrations. ([PR](https://github.com/elastic/integrations/pull/2542))
  
  - 1.9.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2512))
  
  - 1.8.0
     - enhancement: Add routing pipeline to security data_stream, limit to specific providers. ([PR](https://github.com/elastic/integrations/pull/2523))
  
  - 1.7.0
     - enhancement: Expose winlog input language option. ([PR](https://github.com/elastic/integrations/pull/2344))
  
  - 1.6.6
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 1.6.5
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### Traefik - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2513))
  
  - 1.2.2
     - bugfix: Regenerate test files using the new GeoIP database ([PR](https://github.com/elastic/integrations/pull/2339))
  
  - 1.2.1
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### Windows - 1.10.0
Owner: elastic/integrations

Requires: ^7.16.0 || ^8.0.0

Changes since 1.5.0

  - 1.10.0
     - enhancement: Add sysmon event 26 handling ([PR](https://github.com/elastic/integrations/pull/2566))
     - enhancement: Normalise field order and remove event.ingested ([PR](https://github.com/elastic/integrations/pull/2566))
  
  - 1.9.0
     - enhancement: Expose winlog input ignore_older option. ([PR](https://github.com/elastic/integrations/pull/2542))
     - bugfix: Fix preserve original event option ([PR](https://github.com/elastic/integrations/pull/2542))
     - enhancement: Make order of options consistent with other winlog based integrations. ([PR](https://github.com/elastic/integrations/pull/2542))
  
  - 1.8.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2515))
  
  - 1.7.0
     - enhancement: Add provider name check to forwarded/security conditional. ([PR](https://github.com/elastic/integrations/pull/2527))
  
  - 1.6.0
     - enhancement: Expose winlog input language option. ([PR](https://github.com/elastic/integrations/pull/2344))
  
  - 1.5.1
     - bugfix: Change test public IPs to the supported subset ([PR](https://github.com/elastic/integrations/pull/2327))
  



### ZooKeeper Metrics - 1.3.0
Owner: elastic/integrations

Requires: ^7.14.0 || ^8.0.0

Changes since 1.2.0

  - 1.3.0
     - enhancement: Update to ECS 8.0 ([PR](https://github.com/elastic/integrations/pull/2516))
  




To promote these packages use this command:

`elastic-package promote -d=snapshot-production -n -p "apache-1.3.4,aws-1.11.4,aws_logs-0.1.0,azure_metrics-1.0.2,cassandra-1.2.0,cockroachdb-0.2.1,etcd-0.1.1,haproxy-1.1.0,iis-0.8.2,kafka-1.2.0,kubernetes-1.17.0,mysql-1.3.0,nats-1.3.0,nginx-1.3.0,nginx_ingress_controller-1.3.0,prometheus-0.9.0,rabbitmq-1.3.0,redis-1.3.0,stan-1.3.0,system-1.10.0,traefik-1.3.0,windows-1.10.0,zookeeper-1.3.0"`


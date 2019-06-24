# Canary deployments on EKS with Spinnaker


1.Enable/disable canary analysis

hal config canary enable

2.Specify the scope of canary configs
By default, each canary configuration is visible to all pipeline canary stages in all apps. But you can change that so each canary config can be used only within the Spinnaker application in which it was created:

hal config canary edit --show-all-configs-enabled false
Set it to true to revert to global visibility.

3.Set the canary judge
The current default judge is NetflixACAJudge-v1.0. The behavior of this judge is described here (https://www.spinnaker.io/guides/user/canary/judge/).

We use default JUDGE!


4.Identify your metrics provider
hal config canary edit --default-metrics-store STORE
STORE can be…

atlas
datadog
stackdriver
prometheus

hal config canary edit --default-metrics-store prometheus

5.Set up canary analysis for AWS

AWS_ACCESS_KEY_ID=?
hal config canary aws account add canary-account-aws \
    --access-key-id $AWS_ACCESS_KEY_ID \
    --secret-access-key \
    --bucket spinnaker-canary-account-aws-bucket \
    --deployment \
    --no-validate \
    --root-folder kayenta


hal config canary aws edit --s3-enabled true

hal config canary aws canary-account-aws get canary-account-aws

hal config canary aws account list


6.Set up canary analysis to use Prometheus

hal config canary prometheus enable

hal config canary prometheus account add canary-account-prometheus \
    --base-url http://prometheus-server.prometheus.svc.cluster.local

hal config canary prometheus account get canary-account-prometheus



hal config canary aws account edit canary-account-aws \
    --access-key-id AKIAIJD7DNC76UHL3CWA \
    --secret-access-key \
    --bucket spinnaker-canary-account-aws-bucket \
    --region us-west-2 \
    --root-folder kayenta

Note: Try to clear the cache & history in your browser if "canary" feature doesn't show up in the page.



How to use canary analysis:

https://docs.armory.io/spinnaker/configure_kayenta/#enable-canarying-in-application
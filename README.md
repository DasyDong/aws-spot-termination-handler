# AWS Spot Termination Handler

The **Aws Spot Termination Handler** is an operational [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/) built to run on any Kubernetes cluster using AWS [EC2 Spot Instances](https://aws.amazon.com/ec2/spot/). When a user starts the termination handler, the handler watches the AWS [instance metadata service](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html) for [spot instance interruptions](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-interruptions.html) within a customer's account. If a termination notice is received for an instance that’s running on the cluster, the termination handler begins a multi-step cordon and drain process for the node.

You can run the termination handler on any Kubernetes cluster running on AWS, including clusters created with Amazon [Elastic Kubernetes Service](https://docs.aws.amazon.com/eks/latest/userguide/what-is-eks.html).

## What's difference with [aws-node-termination-handler](https://github.com/aws/aws-node-termination-handler/)?
We recommend you follow aws-node-termination-handler repo, it's more professional and scalable.  

And current project is just a lightweight version and more easier to start, is helpful to understand termination logic for everyone 

## How it works

Each `aws-spot-termination-handler` pod polls the notice endpoint until it returns a http status `200`. 
http://169.254.169.254/latest/meta-data/spot/termination-time 

## Getting Started
The termination handler consists of a [ServiceAccount](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/), [ClusterRole](https://kubernetes.io/docs/reference/access-authn-authz/rbac/), [ClusterRoleBinding](https://kubernetes.io/docs/reference/access-authn-authz/rbac/), and a [DaemonSet](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/). All four of these Kubernetes constructs are required for the termination handler to run properly.

### Kubernetes Object Files

We use Kustomize to create a master Kubernetes yaml file .

Read more about Kustomize and Overlays: https://kustomize.io

## Slack Notifications
You are able to setup a Slack incoming web hook in order to send slack notifications to a channel, notifying the users that an instance has been terminated.

Incoming WebHooks require that you set the SLACK_URL environmental variable as part of your PodSpec.

You can also set SLACK_CHANNEL to send message to different slack channel insisted of default slack webhook url's channel.

The URL should look something like: https://hooks.slack.com/services/T67UBFNHQ/B4Q7WQM52/1ctEoFjkjdjwsa22934

Slack Setup:
* Docs: https://api.slack.com/incoming-webhooks
* Setup: https://slack.com/apps/A0F7XDUAZ-incoming-webhooks

```
  env:
    - name: NAMESPACE
      valueFrom:
        fieldRef:
          fieldPath: metadata.namespace
    - name: NODE_NAME
      valueFrom:
        fieldRef:
          fieldPath: spec.nodeName
    - name: POD_NAME
      valueFrom:
        fieldRef:
          fieldPath: metadata.name
    - name: SPOT_POD_IP
      valueFrom:
        fieldRef:
          fieldPath: status.podIP
    - name: SLACK_URL
      value: "https://hooks.slack.com/services/TJH26FK44/BJKR24M1C/GMYZOXmZn6Lg30nl5Hdiz23"
```
## About More motification choices
Refer a good repo [kube-spot-termination-notice-handler](https://github.com/kube-aws/kube-spot-termination-notice-handler), 

##  Contributing
Contributions are welcome!  

## License
This project is licensed under the Apache-2.0 License.

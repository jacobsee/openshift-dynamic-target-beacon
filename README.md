# OpenShift Dynamic Target Beacon

## What is it?

This is an OpenShift beacon for a [Dynamic Target Registration Server](https://github.com/jacobsee/dynamic-target-registration-server). It is intended to run as a deployment on an OpenShift cluster, and must run as a service account bound to the `openshift-cluster-monitoring-view` ClusterRole in order to register to the server with proper credentials.

## How is it configured?

| Variable Name | Description | Example |
| --- | --- | --- |
| `SERVER_URL` | A fully-formed URL (with protocol prefix) at which the registration server can be reached | `https://1.2.3.4` |
| `CLUSTER_URL` | The base URL of this cluster to register for monitoring, without `apps` or `api` prefixes | `my-cluster.my-domain.com` |
| `AUTH_TOKEN` | The token to pass in an `Authorization` header to the registration server | `12345abcde` |
| `DEBUG` | If set, will not attempt to fetch an API token using the incluster service account and will pass a dummy API token to the registration service instead (useful for local debugging, leave unset for default behavior) | `1` |

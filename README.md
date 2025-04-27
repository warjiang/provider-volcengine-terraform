# Provider VolcengineTerraform

`provider-volcengine-terraform` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
VolcengineTerraform API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/warjiang/provider-volcengine-terraform):
```
up ctp provider install warjiang/provider-volcengine-terraform:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1alpha1
kind: ControllerConfig
metadata:
  name: volcengine-terraform-config
  labels:
    app: provider-volcengine-terraform
spec:
  args: ["-d"]
  image: ghcr.io/warjiang/provider-volcengine-terraform:v0.1.4
---
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-volcengine-terraform
spec:
  package: xpkg.upbound.io/warjiang/provider-volcengine-terraform:v0.1.1
  controllerConfigRef:
    name: volcengine-terraform-config

EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/warjiang/provider-volcengine-terraform).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

Build & Push xpkg:
```console
crossplane xpkg build --package-root=package/
crossplane xpkg login --token={your upbound token}
crossplane xpkg push xpkg.upbound.io/warjiang/provider-volcengine-terraform:v0.1.0 -f ./package/provider-volcengine-terraform-684452faf24b.xpkg
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/warjiang/provider-volcengine-terraform/issues).

# Kubectl set-namespace plugin

This plugin will allow you to change the default namespace in the current context in the kubeconfig file.

### Requirements

kubectl v1.12.0 or later

### Installing on macOS

```
 curl -L https://github.com/FrankFuenmayor/kubectl-kubeconfig-plugins/releases/download/0.0.1/kubectl-kubeconfig-plugins_darwin_amd64.tar | tar xvz && mv kubectl-* /usr/local/bin
```

### Download

Go to [realease page](https://github.com/FrankFuenmayor/kubectl-kubeconfig-plugins/releases/latest) and download the 
file`kubectl-kubeconfig-plugins_darwin_amd64.tar` extract it and move them to a folder in your path.  

# Available plugins


### set-namespace

```bash
kubectl set-namespace myapp
```

Changes the current context's default namespace to `myapp`


### aws-update-kubeconfig

```bash
kubectl aws-update-kubeconfig
```

Update kubeconfig file with the clusters information from AWS

**Note**: This command requires aws-cli installed and configured to work









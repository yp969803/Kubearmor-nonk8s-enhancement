# Kubearmor Non-k8s Enhancement

## Setup KubeArmor in Unorchestrated mode on a BPF LSM node

- Download kubearmor package
```
curl https://github.com/kubearmor/KubeArmor/releases/download/v1.4.0/kubearmor_1.4.0_linux-amd64.deb
```
- Install kubearmor package
```
sudo apt --no-install-recommends install ./kubearmor_1.4.0_linux-amd64.deb
``` 
- Start kubearmor
```
sudo systemctl start kubearmor
``` 
- Install karmor cli
```
curl -sfL http://get.kubearmor.io/ | sudo sh -s -- -b /usr/local/bin
```


## Policy

```
curl -O https://github.com/kubearmor/KubeArmor/raw/main/examples/kubearmor_containerpolicy.yaml
```

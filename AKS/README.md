#### AKS

Persistent volumes on Kubernetes via AKS

Reference:
`https://docs.microsoft.com/en-us/azure/aks/azure-disks-dynamic-pv`

#### Built-in storage classes
`kubectl get sc`

Example output:

```
NAME                PROVISIONER                AGE
default (default)   kubernetes.io/azure-disk   1h
managed-premium     kubernetes.io/azure-disk   1h
```

#### Verify the namespace ####
Before creating a PVC. Make sure that the PVC will be created on the same namespace as the pods

To list all the namespaces:
`kubectl get namespaces`

To see which name space is used. Run the following command (empty = defualt namespace):
`kubectl config view | grep namespace`

To switch to a different namespace use the command below:
`kubectl config set-context --current --namespace=<myNamespace>`

#### Create a persistent volume claim (PVC)
Persistent volume claim (PVC) dynamically allocates a persistent volume (PV)
This is how it looks like:
Pod -> PVC -> PC -> Host machine

PVC and PC provide an abstraction of the application from the specific storage type.

apply the `storage.yaml` file to make the storage

Apply the storage:
`kubectl apply -f storage.yaml`

*Note the status of the storage will be pending until a pod will use it.

#### Inspect persistent volume (pvc)
`kubectl get pvc`

#### Use the persistent volume
Create a producer pod. Use the `producer_pod.yaml` file.

Run the following to apply:
`kubectl apply -f producer_pod.yaml`

Create a consumer pod. Use the `consumer_pod.yaml` file.

Run the following to apply:
`kubectl apply -f producer_pod.yaml`

You now have a running pod with your Azure disk mounted in the `/app` directory.
This configuration can be seen when inspecting your pod via `kubectl describe pod mypod`.

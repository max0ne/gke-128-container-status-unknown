
This repo reproduces `ContainerStatusUnknown` issue observed in GKE 1.28.

Usage:

1. Build docker image, push to a container registry
2. Update `image` in `k8s.yaml`
3. `kubectl apply -f k8s.yaml`

It should run a pod that:
1. self destructs once every 10 seconds, by using more ephemeral storage than it requests.
2. responds to SIGTERM, but does not exit within `terminationGracePeriodSeconds`

In a normal case, pod should be evicted by kubelet for exceeding its ephemeral storage limit and be forcefully killed after `terminationGracePeriodSeconds`, see `good_pod.yaml` as an example.

However at a small probability, after `terminationGracePeriodSeconds`, the container gets into `ContainerStatusUnknown` state, see `bad_pod.yaml` as an example.

```
kubectl get po

NAME                                                READY   STATUS                   RESTARTS   AGE
gke-128-container-status-unknown-7955bdbc96-2kxmd   0/1     Error                    0          5m43s
gke-128-container-status-unknown-7955bdbc96-4wqjr   0/1     Error                    0          3m16s
gke-128-container-status-unknown-7955bdbc96-55xm4   0/1     Error                    0          2m7s
gke-128-container-status-unknown-7955bdbc96-5n4np   0/1     Error                    0          8m33s
gke-128-container-status-unknown-7955bdbc96-6vgmv   0/1     Error                    0          92s
gke-128-container-status-unknown-7955bdbc96-75pnh   0/1     Error                    0          7m59s
gke-128-container-status-unknown-7955bdbc96-7gv9r   0/1     Error                    0          14m
gke-128-container-status-unknown-7955bdbc96-8pzpb   0/1     Error                    0          6m51s
gke-128-container-status-unknown-7955bdbc96-9zms8   0/1     ContainerStatusUnknown   1          3m59s
```

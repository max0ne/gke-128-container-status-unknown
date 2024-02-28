
This repo reproduces `ContainerStatusUnknown` issue observed in GKE 1.28.

Usage:

1. Build docker image, push to a container registry
2. Update `image` in `k8s.yaml`
3. `kubectl apply -f k8s.yaml`

It should run a pod that:
1. self destructs once every 10 seconds, by using more ephemeral storage than it requests.
2. responds to SIGTERM, but does not exit within `terminationGracePeriodSeconds`

In a normal case, pod should be evicted by kubelet for exceeding its ephemeral storage limit and be forcefully killed after `terminationGracePeriodSeconds`, see `good_pod.yaml` as an example.

However at some probability, after `terminationGracePeriodSeconds`, the container gets into `ContainerStatusUnknown` state, see `bad_pod.yaml` as an example.

```
kubectl get po

NAME                                                READY   STATUS                   RESTARTS   AGE
NAME                                                READY   STATUS                   RESTARTS   AGE
gke-128-container-status-unknown-7955bdbc96-2j6pz   0/1     Error                    0          5m12s
gke-128-container-status-unknown-7955bdbc96-2kxmd   0/1     Error                    0          16m
gke-128-container-status-unknown-7955bdbc96-4wqjr   0/1     Error                    0          13m
gke-128-container-status-unknown-7955bdbc96-52srf   0/1     Error                    0          8m
gke-128-container-status-unknown-7955bdbc96-55xm4   0/1     Error                    0          12m
gke-128-container-status-unknown-7955bdbc96-5n4np   0/1     Error                    0          18m
gke-128-container-status-unknown-7955bdbc96-5rtc4   0/1     ContainerStatusUnknown   1          61s
gke-128-container-status-unknown-7955bdbc96-6q5nn   0/1     Error                    0          106s
gke-128-container-status-unknown-7955bdbc96-6vgmv   0/1     Error                    0          11m
gke-128-container-status-unknown-7955bdbc96-75pnh   0/1     Error                    0          18m
gke-128-container-status-unknown-7955bdbc96-7gv9r   0/1     Error                    0          24m
gke-128-container-status-unknown-7955bdbc96-8pzpb   0/1     Error                    0          17m
gke-128-container-status-unknown-7955bdbc96-9zms8   0/1     ContainerStatusUnknown   1          14m
gke-128-container-status-unknown-7955bdbc96-b5pcq   0/1     Error                    0          11m
gke-128-container-status-unknown-7955bdbc96-blfgs   0/1     Error                    0          14m
gke-128-container-status-unknown-7955bdbc96-c84r8   1/1     Running                  0          27s
gke-128-container-status-unknown-7955bdbc96-clw7h   0/1     Error                    0          6m18s
gke-128-container-status-unknown-7955bdbc96-crhxv   0/1     ContainerStatusUnknown   1          15m
gke-128-container-status-unknown-7955bdbc96-df96g   0/1     Error                    0          2m55s
gke-128-container-status-unknown-7955bdbc96-dfqbg   0/1     ContainerStatusUnknown   1          22m
gke-128-container-status-unknown-7955bdbc96-dqt7h   0/1     ContainerStatusUnknown   1          9m41s
gke-128-container-status-unknown-7955bdbc96-hrdxw   0/1     Error                    0          21m
gke-128-container-status-unknown-7955bdbc96-jd6fh   0/1     Error                    0          5m43s
gke-128-container-status-unknown-7955bdbc96-jzbf8   0/1     Error                    0          9m8s
gke-128-container-status-unknown-7955bdbc96-kcnkt   0/1     ContainerStatusUnknown   1          13m
gke-128-container-status-unknown-7955bdbc96-ksfgg   0/1     Error                    0          4m38s
gke-128-container-status-unknown-7955bdbc96-lh55v   0/1     ContainerStatusUnknown   1          8m33s
gke-128-container-status-unknown-7955bdbc96-q9btd   0/1     Error                    0          20m
gke-128-container-status-unknown-7955bdbc96-qcskz   0/1     ContainerStatusUnknown   1          10m
gke-128-container-status-unknown-7955bdbc96-rrvfn   0/1     Error                    0          16m
gke-128-container-status-unknown-7955bdbc96-snvw2   0/1     Error                    0          10m
gke-128-container-status-unknown-7955bdbc96-tnhm7   0/1     Error                    0          23m
gke-128-container-status-unknown-7955bdbc96-vnt62   0/1     Error                    0          17m
gke-128-container-status-unknown-7955bdbc96-ws47z   0/1     Error                    0          4m3s
gke-128-container-status-unknown-7955bdbc96-x46pq   0/1     Error                    0          20m
gke-128-container-status-unknown-7955bdbc96-xg5lq   0/1     Error                    0          2m21s
gke-128-container-status-unknown-7955bdbc96-xrrmx   0/1     ContainerStatusUnknown   1          6m51s
gke-128-container-status-unknown-7955bdbc96-z2dzc   0/1     Error                    0          3m29s
gke-128-container-status-unknown-7955bdbc96-z4xjp   0/1     Error                    0          19m
gke-128-container-status-unknown-7955bdbc96-z7jr9   0/1     ContainerStatusUnknown   1          21m
gke-128-container-status-unknown-7955bdbc96-zs2pn   0/1     Error                    0          7m26s
```

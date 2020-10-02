# workaround-openshift-rt

Currently there's an issue with the routing table when running OKD 4/OpenShift 4 on Azure.

## Issue:
- TLS handshake errors occur although TCP communication is possible. A traffic capture shows large packets being discarded (see "Diagnostic Steps").
- Unexpected ICMP fragmentation needed messages are received for direct communications happening between OpenShift nodes but without vxlan encapsulation. Requested MTU is lower than the one set in both ends and/or required by any intermediate element.
- Routing cache shows bad entries as described in "Diagnostic Steps".
- After some time, the OpenShift Cluster becomes very slow and many operators start to become unhealthy (degraded state).

More infos:
https://access.redhat.com/solutions/5252831

## Workaround

I made a little go program which runs every 30 minutes on all masters and nodes and flushes the routing table cache.
The program has to run privileged and as a daemonset.

### Steps to get it running

1. Build the program and push it to your registry.
2. Once it's up there change the image url in the daemonset.
3. oc new-project workaround-openshift-rt && oc project workaround-openshift-rt
4. oc create sa workaround-openshift-rt
5. oc adm policy add-scc-to-user privileged system:serviceaccount:workaround-rt:workaround-openshift-rt
6. oc create -f daemonset.yaml

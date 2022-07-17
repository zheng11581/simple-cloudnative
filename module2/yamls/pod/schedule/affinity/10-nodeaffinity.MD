### Which Node the pod should be scheduled?

1. nodeAffinity

- The schedulable nodes `required` label `kubernetes.io/os` `In` `linux` (all nodes)

2. nodeAntiAffinity

- The schedulable nodes anti-`preferred` label `disktype` `In` `ssd` with `weight=1` (node2)
- The schedulable nodes anti-`preferred` label `disktype` `In` `hdd` with `weight=50` (node1)

3. nodeAffinity like all nodes, nodeAntiAffinity dislike node1 more

4. Pod will scheduled to node2

### Which Node the pod should be scheduled?

1. nodeAffinity

- The schedulable nodes `required` label `kubernetes.io/os` `In` `linux` (all nodes)
- The schedulable nodes `preferred` label `disktype` `In` `ssd` with `weight=1` (node2)
- The schedulable nodes `preferred` label `disktype` `In` `hdd` with `weight=50` (node1)

3. All nodes fit required rules, node1 is more fit preferred rules
4. Pod will scheduled to node1

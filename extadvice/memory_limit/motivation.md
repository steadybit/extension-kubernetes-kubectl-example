Your Kubernetes pods of ${target.steadybit.label} can use more memory than defined in ```request``` as Kubernetes uses this only for scheduling pods.
Hence, you should configure an upper limit to prevent using the entire memory of the node at cost of other pods.
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/fagneracr/queue/master)

# Go-Q

Just a queue system in golang, managing tasks or requests by organizing them in a first-in, first-out (FIFO) order.




## CreateQ

The `Createq` function creates a new queue in the `System`. It takes a `ConfigQueue` struct as input, which contains parameters for configuring the new queue. The function performs the following steps:

- Checks that the name of the queue is not empty and that the queue does not already exist in the `System`.
- Creates a new `queueConf` struct with the provided configuration parameters, generates a new UUID for the queue, and saves the configuration to a file.
- Creates the necessary directory structure for the queue.
- Adds the new queue configuration to the `e.Queue` slice in the `System` struct and returns an error if there was any problem during the process.

## ListQ

The `ListQ` method takes a variadic number of string arguments which represent the names of queues. It returns a slice of `ConfigQueue` structs which represent the configuration of the specified queues or all queues in the `System`. The method performs the following steps:

- If one or more queue names are provided, the method iterates over each name and appends the corresponding `ConfigQueue` to the output slice. If a `ConfigQueue` for a provided queue name is not found, it is skipped.
- If no queue names are provided, the configuration of all queues in the `System` is appended to the output slice.
- The `produceOutConfig` function is called to convert a `queueConf` struct to a `ConfigQueue` struct, which is then appended to the output slice.

## DeleteQ

The `DeleteQ` method takes a string argument which represents the name of the queue to be deleted from the `System`s It returns an error if the queue does not exist in the `System` or if there was an error deleting the queue's struct and config files. The method performs the following steps:

- Checks if the queue with the provided name exists in the `System` using the `Exists` method.
- If the queue exists, obtains a lock on the `System`'s mutex to ensure exclusive access to the Queue slice.
- Iterates over each `Queue` in the `System` and finds the Queue with a matching name.
- Deletes the `Queue`'s struct and config files using the deleteStructureQ function and removes the Queue from the `System`'s `Queue` slice.
- Returns an error if there was an error deleting the queue or if no `Queue` with a matching name is found.
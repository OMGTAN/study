from rediscluster import RedisCluster

startup_nodes = [{"host": "127.0.0.1", "port": "7000"}]

rc = RedisCluster(
    startup_nodes=startup_nodes,
    decode_responses=True,
    host_port_remap=[
        {
            'from_host': '172.18.0.2',
            'from_port': 7000,
            'to_host': 'localhost',
            'to_port': 7000,
        },
        {
            'from_host': '172.22.0.1',
            'from_port': 7000,
            'to_host': 'localhost',
            'to_port': 7000,
        },
    ]
)

## Debug output to show the client config/setup after client has been initialized.
## It should point to localhost:7000 for those nodes.
print(rc.connection_pool.nodes.nodes)

## Test the client that it can still send and recieve data from the nodes after the remap has been done
print(rc.set('foo', 'bar'))
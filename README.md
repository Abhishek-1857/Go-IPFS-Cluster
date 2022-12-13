# Go-IPFS-Cluster
A IPFS cluster(Adding and reteiving data) in go
Setting up an IPFS private node locally:- <br>
Install the prereqs from https://medium.com/@s_van_laar/deploy-a-private-ipfs-network-on-ubuntu-in-5-steps-5aad95f7261b <br> 
run the following commands <br>
1. IPFS_PATH=~/.ipfs ipfs init <br>
Remove all pre-existing bootstrap node for the network <br>
2. IPFS_PATH=~/.ipfs ipfs bootstrap rm --all <br>
Now, change the configuration API and Gateways inside the config file inside the ".ipfs" folder <br>
On the bootstrap node, create a swarm key by using <br>
3. echo -e "/key/swarm/psk/1.0.0/\n/base16/\n`tr -dc 'a-f0-9' < /dev/urandom | head -c64`" > ~/.ipfs/swarm.key <br>
and then send it to the clients you wish to share the network with <br>
Find the peerID for the bootstrap node by using <br>
4. IPFS_PATH=~/.ipfs ipfs config show | grep "PeerID" <br>
Now, add bootstrap node to client node config usring <br>
5. IPFS_PATH=~/.ipfs ipfs bootstrap add /ip4/<ip address of bootnode>/tcp/4001/ipfs/<peer identity hash of bootnode> <br>
6. export LIBP2P_FORCE_PNET=1<br>
7. IPFS_PATH=~/.ipfs ipfs daemon <br>

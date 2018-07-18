## KNS Explained
KNS is the Kowala Name Service, a distributed, open, and extensible naming system based on the Kowala blockchain.

KNS can be used to resolve Kowala addresses, from human-readable form like “kowala.tech” into machine-readable identifiers, including Kowala addresses.

KNS is very similar to DNS, the Internet’s Domain Name Service, but has significantly different architecture, due to the capabilities and constraints provided by the Kowala blockchain. Like DNS, KNS operates on a system of dot-separated hierarchial names called domains, with the owner of a domain having full control over the distribution of subdomains.

First component of the KNS is **registry**. The registry is a central directory of KNS where all of the KNS domains are kept.
Every name in KNS can be found by looking it up in the KNS registry, and it’s the only component you need the address for.

**Registrars** are the second component of KNS, and are responsible for allocating new names to users. Registrars don’t have any special permissions — they just use their ability to tell the registry to create subdomains.

**Resolvers** are contracts that can tell you the resource associated with a name — such as an Kowala address.


## Namehash
KNS does not operate directly on names — there are a number of reasons for this, one of which is that parsing and processing text on the blockchain is very inefficient. Instead, KNS operates on secure hashes.

KNS hashes names using a system called namehash. Here’s the full definition of namehash:
* namehash('') -> '0x0000000000000000000000000000000000000000000000000000000000000000'
* namehash('a.xyz') -> sha3(namehash('xyz'), sha3('a'))


## How to use KNS?
To use KNS you will need to deploy all the components to the network.

Deploying KNS with *kowala.tech* domain and assign contract under this domain will look something like this (truffle test pseudo-code):
1. Deploy KNS Registry
	`KNS.new();`
2. Deploy FIFSRegistrar with parameters 
	1. **nsAddr** The address of the KNS registry.
	2. **node** The node that this registrar administers.

In our example it will be like this
`FIFSRegistrar.new(this.kns.address, namehash('tech'));`

3. Deploy Resolver with one parameter
	1. **nsAddr** The address of the KNS registry.

`PublicResolver.new(this.kns.address);`

4. Add root domain to the kns registry
	`kns.setSubnodeOwner(0, web3.sha3('tech'), this.registrar.address);`
    * First parameter: **node** The parent node.
    * Second parameter: **label** The hash of the label specifying the subnode.
    * Third parameter: **owner** The address of the new owner.

5. Register *kowala* domain under *.tech* root domain.
	`registrar.register(web3.sha3('kowala'), accounts[0]);`
    * First parameter: **subnode** The hash of the label to register.
    * Second parameter: **owner** The address of the new owner.

6. When we have a new domain under root domain, we should add resolver to that domain.
	`kns.setResolver(namehash('kowala.tech'), resolver.address);`
	* First parameter: **node** The node to update.
	* Second parameter: **resolver** The address of the resolver.
7. Now we can use our resolver to set our domain to point to address of a contract.
	`resolver.setAddr(namehash('kowala.tech'), kowalaContract.address);`
	* First parameter: **node** The node to update.
	* Second parameter: **addr** The address to set.

8. Having set everything up, we can now use simple function from our resolver to translate domain name to an address
	`resolver.addr(namehash('kowala.tech'));`

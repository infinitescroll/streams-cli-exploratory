# A sample of streams-core

### How archival works
1. The stream-owner creates a stream, which under the hood:
  - Instantiates a new thread + db
  - Instantiates an FFS instance (with a new wallet).
  - creates a number of collections:
    - OwnerPointer 
    - Resource
    - StreamsMetadata
    - StreamsConfiguration (could be the same as metadata?)
    - Deals
  - Creates an OwnerPointer instance, pointing to the stream-owner
  - Creates a StreamsMetadata instance
  - Creates a StreamsConfiguration instance (with some defaults)
2. The stream-owner's wallet needs to get funded for deal making (how?)
3. The stream-owner adds a resource, like a GitHub repository.
3. Once the wallet is funded, a new StreamsArchiver service gets instantiated, passing the StreamsConfiguration. The archiver will:
  - At every interval (interval determined by the StreamConfig), the archiver will loop through the Stream's Resource collection, and get information about each resource, like it's API endpoint to fetch data and the type of resource it is.
  - The archiver will look to the stream-owner thread+db, gathering the associated API token if necessary.
  - With that information, the archiver will fetch the latest information from the data source (the app itself), using the tokens from the stream-owner thread.
  - The archival service will take the data, and formulate it into a deal within given parameters (the parameters would be determined by the Streams Config, which is determined by FFS config).
  - The archival service watches the deal
  - Once the deal is active, the streams-archiver writes the deal metadata to the stream-thread, Deal collection.

### Building consumer apps (winamp skins)
Let's imagine we want to build a web app that shows what's being archived (i.e., what data is being turned into deals via FFS and put on IFPS), while also building robust search on top of that archive. How could we achieve that?

So essentially the app we're building is:

- a search engine built on top of streams
- a web app to provide a UI for the search

For simplicity, let's pretend that all streams data is public, so anyone can access it (streams core v0.0.1). The consumer app could do the following:

1. boot up
2. get a list of all streams (in reality, only getting all _public_ streams, or private streams that have granted access to this app)
3. register a [Listener](https://godoc.org/github.com/textileio/go-threads/api/client#Client.Listen) on each of the streams to know when changes to the streams Resource and Deal collections occur
4.  When a new resource gets added, the consumer application factors that in to its own backend. In our case, this might register a webhook or event polling mechanism to the resource.
5. When a new deal gets created, show some informative graphics in the UI.

QUESTION: How does this app gain access to each stream-owner's resource tokens?

### Future of privacy and app-siloed data
In the stream-metadata, we can provide different types of streams. To start, we could mark all as "public" and thus, consumer applications automatically get access to the data inside those streams (i.e., no one needs to give permission). 

Later, streams can have more granular privacy types, and the consumer apps would need to request permission to these streams to be able to view/use the data. Those streams can also get created in an electron app on the users own machine. So they can have full control over the thread and database and who is getting invited to it. 

### What's implemented in this exploratory repo
The bullet points might be implemented in code in a different order:

1. Create an org: <br />
`streams-cli org create [orgname]`<br />
`streams-cli org create openworklabs`<br />
- Creates a new thread for the organization
- Creates an OwnerMetadata collection
- Creates an OwnerMetadata instance (this is where org access tokens, and stream-pointers get stored)
- Creates a new StreamPointer collection on the org thread (for future use, when creating streams that belong to this org)
- Creates an OwnerPointer record on the streams-master thread, pointing at the threadID of the newly created org

2. Create a stream, passing the org from step 1:<br />
`streams-cli stream create [streamname] [ownername] [ownertype]`<br />
`streams-cli stream create stream-name-1 openworklabs organization`
- Creates a new thread for the stream
- Creates a new FFS instance for the stream
- Creates a new StreamsMeta collection for storing stream metadata
- Creates a new StreamsMeta instance for storing information about the FFS instance and wallet (later this could include streams settings too)
- Creates a new StreamPointer on the org's thread

### Not implemented:
- stream-meta thread of pointers to all the streams
- pointers from stream back to its owner
- Single stream functionality

### Open questions:
- [ ] Where is the best place to store access tokens? Do apps need these access tokens? How can users share their archived data and resources without sharing access tokens? One thing to think about here is that if we put the accesss tokens in the stream-owner thread, then all streams essentially share the same access token. Putting the access token at the stream level itself enables more granular privacy and permission management, but it creates a more complex streams-creation expreience (you have to create a new access token for each stream or grab it from another stream that already has it). It also makes it harder to share user data with a consumer app, without sharing the access tokens. 
- [ ] Is it safe to create a DB and a collection that already exist? Can we ignore those "____ already exists" errors?
- [ ] Authentication for streams-cli
- [ ] Security with powergate and partitioning a single node's wallet to use many wallets (1 per stream)
- [ ] Handling edits / removal of data, do we fork a thread? 
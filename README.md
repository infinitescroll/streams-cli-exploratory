What's implemented so far (the bullet points might be implemented in code in a different order):

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

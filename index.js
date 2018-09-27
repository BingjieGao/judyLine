const driver = require('bigchaindb-driver')

// BigchainDB server instance (e.g. https://test.bigchaindb.com/api/v1/)
const API_PATH = 'https://test.bigchaindb.com/api/v1/'

// Create a new keypair.

const alice = new driver.Ed25519Keypair()
console.log(alice.privateKey)
console.log(alice.publicKey)
const privateKey = "x1RkdikfnaUqCbYtroYRkuL64GZdU6twSLV6BNehsvE"
const publicKey = "D1ACjwzXEyP8s5Yjs6xDhfhd3yM7d92jxnAznysDFC3D"

// postData from argv[2]
const postData = process.argv[2]

// Construct a transaction payload
const tx = driver.Transaction.makeCreateTransaction(
    // Define the asset to store, in this example it is the current temperature
    // (in Celsius) for the city of Berlin.
    JSON.parse(postData),

    // Metadata contains information about the transaction itself
    // (can be `null` if not needed)
    { what: 'My second BigchainDB transaction' },

    // A transaction needs an output
    [ driver.Transaction.makeOutput(
            driver.Transaction.makeEd25519Condition(publicKey))
    ],
    publicKey
)

// Sign the transaction with private keys
const txSigned = driver.Transaction.signTransaction(tx, privateKey)

console.log(`txSigned is ${JSON.stringify(txSigned)}`)

// Send the transaction off to BigchainDB
const conn = new driver.Connection(API_PATH, {
  app_id: "82826d8b",
  app_key: "8ea557d8f236db15626c7d68a04eca6b"
})

conn.postTransactionCommit(txSigned)
    .then(retrievedTx => console.log('Transaction', retrievedTx.id, 'successfully posted.'))
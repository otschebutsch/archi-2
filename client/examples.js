const computers = require('./computers/client');

const client = computers.Client('http://localhost:8080');

const readline = require('readline').createInterface({
  input: process.stdin,
  output: process.stdout
})

// Scenario 1: Display available computers.
client.listComputers()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available computers:');
        console.dir(list);
    })
    .catch((e) => {
        console.log(`Problem listing available computers: ${e.message}`);
    });


// Scenario 2: Attach one of the discs to one of the computers
  client.setDisc()
	.then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Attaching one of the discs to one of the computers');
		console.log('Done');
    console.dir(resp);

    })
    .catch((e) => {
        console.log(`Problem creating a new computer: ${e.message}`);
    });




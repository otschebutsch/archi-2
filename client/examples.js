const channels = require('./computers/client');

const client = channels.Client('http://localhost:8080');

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
        console.log(`Problem listing available channels: ${e.message}`);
    });



  client.setDisc()
	.then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create computer response:');
		console.log('Done');
    console.dir(resp);

    })
    .catch((e) => {
        console.log(`Problem creating a new channel: ${e.message}`);
    });





/*
// Scenario 1: Display available computers.
client.listComputers()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available computers:');
        Array.prototype.forEach.call(list, (c) => console.log(c.name));
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });

// Scenario 2: Create new channel.
client.createComputer('my-new-channel')
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create computer response:', resp);
        return client.listComputers()
            .then((list) => Array.prototype.map.call(list, (c) => c.name).join(', '))
            .then((str) => {
                console.log(`Current channels: ${str}`);
            })
    })
    .catch((e) => {
        console.log(`Problem creating a new channel: ${e.message}`);
    });

client.listComputers()
    .then((list) => Array.prototype.filter.call(list, (computer) => computer.id === 1))
    .then((elem) => console.log(elem))
    .catch((e) => {
        console.log(`Problem listing available channels2: ${e.message}`);
    });

client.createDisk(1, 500)
    .then((resp) => {
        console.log('=== Scenario 3 ===');
        console.log('Create disc response:', resp);
        return client.listComputers()
            .then((list) => Array.prototype.filter.call(list, (computer) => computer.id === 1))
            .then((elem) => console.log(elem));
    });
	*/

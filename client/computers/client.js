const http = require('../common/http');


const Client = (baseUrl) => {

    const client = http.Client(baseUrl);
	return {
        listComputers: () => client.get('/computers'),
		setDisc: () => client.get('/computers'),
    }

};

module.exports = { Client };





/*
const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listComputers: () => client.get('/computers'),
        createComputer: (name, cpuCount) => client.post('/computers', { name, cpuCount }),
        createDisk: (idComputer, space) => client.post('/disks', { idComputer, space }),
        listDisks: () => client.get ('/disks')
    }

};

module.exports = { Client };
*/

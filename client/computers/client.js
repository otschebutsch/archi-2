const http = require('../common/http');


const Client = (baseUrl) => {

    const client = http.Client(baseUrl);
	return {
        listComputers: () => client.get('/computers'),
		setDisc: () => client.get('/computers'),
    }

};

module.exports = { Client };


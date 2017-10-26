var lambda = require('./index');

lambda.handler(null, null, function(error, text){
	console.log(text);
});
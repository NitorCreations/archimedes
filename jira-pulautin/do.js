var JiraClient = require('jira-connector');
 
var jira = new JiraClient( {
	host: 'nitor-archimedes.atlassian.net',
	basic_auth: {
        username: "thedude@mailinator.com",
		password: "Tehdude2017#"
    }
});

jira.issue.createIssue(
{"fields": {
        "project": {
            "id": "10001"
        },
        "issuetype": {
            "id": "10000"
        },
        "summary": "something's wrong",
        "customfield_10005": "foobarz",
        "reporter": {
            "name": "thedude" //"thedude@mailinator.com"
        }
	}
}
);
 
jira.issue.getIssue({
    issueKey: 'DEMO-6'
}, function(error, issue) {
	if(error){
    	console.log('Error: ' + error);
	}else{
	    console.log('Status: ' + issue.fields.status.name);
	}
});
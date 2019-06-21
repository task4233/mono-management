var ghpages = require("gh-pages");

ghpages.publish("./dist", {
		branch: "gh-pages",
		user: {
			name: "Kaoru Chisen",
			email: "cordx56@cordx.net"
		},
		message: "Web page AUTO DEPLOY [ci skip]"
	}, function(err) {
		if (err) {
			console.log(err);
			process.exit(1);
		} else {
			console.log("Successfully Updated!");
		}
	});

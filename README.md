## Questions API
built on echo frame work
uses postgresql
runs on port 8181

to run run/build/install `main.go`

	/questions
		POST	insert new question
		`{
			text:"what is the answer"
			answers:[
				"it could be this",
				"or  this",
				"and this"
			]
			correct:3
		}`
	/question/:id	
		GET 	get question by id
		PUT 	update question by id

	/answers
		POST	`[
				{
					id:"id of question"
					correct:3
				},
				{
					id:"id of question"
					correct:3
				}

			]`

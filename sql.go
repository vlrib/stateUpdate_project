package main

//Project Queries
const CreateProject = `
	INSERT INTO projects (id, name, status)
	VALUES(:id, :name, :status)
`

const SelectProjectByName = `
	SELECT id, name, status FROM projects WHERE name=$1
`

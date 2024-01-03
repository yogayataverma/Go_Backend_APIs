These are Restful APIs in Go Language.

To Run this project:
1. Clone Repo.
2. OPen folder in VS Code.
3. Download go lang and sqlite db browser on system.
4. Import the sqlite db file in it.
5. Download Postman
6. Use the follwing urls in Postman:
   1. To get all tasks: localhost:8080/tasks(GET)
   2. To get particular task: localhost:8080/task/1(GET)
   3. To insert task:  localhost:8080/tasks(POST)
   4. To update task:  localhost:8080/tasks/1(PUT)
   5. To delete task:  localhost:8080/tasks/1(DELETE)


While inserting new task use this reference:
	ID          int    
	Title       string 
	Description string 
	Due_Date    string 
	Status      string 

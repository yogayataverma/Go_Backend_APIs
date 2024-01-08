# Go Language RESTful APIs with SQLite

## To Run the Project:

1. **Clone the Repository:**
   - Clone this repository to your local machine.

2. **Open Folder in VS Code:**
   - Open the cloned folder in Visual Studio Code.

3. **Download Go Lang:**
   - Download and install Go from [here](https://golang.org/dl/).

4. **Download SQLite DB Browser:**
   - Download and install SQLite DB Browser from [here](https://sqlitebrowser.org/dl/).

5. **Import SQLite DB File:**
   - Open SQLite DB Browser.
   - Import the provided SQLite DB file into the browser.

6. **Run the Project:**
   - Open the terminal in VS Code.
   - Run the following command to start the Go server:
     ```bash
     go run main.go
     ```

7. **Download Postman:**
   - Download and install Postman from [here](https://www.postman.com/downloads/).

8. **Use Postman for API Testing:**
   - Open Postman.
   - Use the following URLs for API testing:
     - To get all tasks: `localhost:8080/tasks` (GET)
     - To get a particular task: `localhost:8080/task/1` (GET)
     - To insert a task: `localhost:8080/tasks` (POST)
     - To update a task: `localhost:8080/tasks/1` (PUT)
     - To delete a task: `localhost:8080/tasks/1` (DELETE)
   - Use the provided reference for inserting a new task: 
     ```json
     {
       "ID": 1,
       "Title": "Sample Task",
       "Description": "Description of the task",
       "Due_Date": "2022-12-31",
       "Status": "Pending"
     }
     ```

## Additional Notes:

- Customize the Go code and SQLite database as needed.

Happy API testing with Go and SQLite! 🚀

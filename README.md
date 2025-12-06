# TODO-CLI BACKEND

## Entities
*Category*
    *Properties:*
        - Title
        - Color
    *Behavior:*
        - Create a new Category
        - List user categories with progress status
        - Edit a category

*Task*
    *Properties:*
        - Title
        - DueDate
        - Category
        - IsDone
    *Behavior:*
        - Create a new task
        - List user today task
        - List user tasks by date
        - change task status (done/undone)
        - Edit a task

*User*
    *Properties:*
        - Id
        - Email
        - Password
    *Behavior:*
        - Register a user
        - Log-in user

----------------------------------------------------------

*User Story (use-cases)*
    - User should be registered
    - User should be able to log in to the application
    - User can create a new Category
    - User can create a new Task
    - User can see the list of categories with progress status
    - User can see the Today's tasks
    - User can see the Tasks by date
    - User can Done/Undone a task
    - User can edit a task
    - User can edit a category

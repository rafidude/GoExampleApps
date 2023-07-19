const form = document.getElementById('todoForm')
const saveButton = document.getElementById('saveButton')

form.addEventListener('submit', async (event) => {
    event.preventDefault()

    const id = form.getAttribute('data-id')
    const title = document.getElementById('title').value
    const completed = document.getElementById('completed').checked

    // Construct form data
    const formData = new URLSearchParams()
    formData.append('title', title)
    formData.append('completed', completed)

    try {
        let response
        let todo

        if (id) {
            // If the form has a 'data-id' attribute, update the existing todo
            response = await fetch(`/todos/${id}`, {
                method: 'POST',
                body: formData,
            })

            if (!response.ok) {
                throw new Error('Failed to update todo')
            }

            todo = await response.json()

            // Update the row in the table
            const row = document.querySelector(`.todo-row[data-id="${id}"]`)
            row.querySelector('.todo-title').innerText = todo.title
            row.querySelector('.todo-completed').innerText = todo.completed
                ? 'true'
                : 'false'

            form.removeAttribute('data-id') // Remove the 'data-id' attribute after updating
        } else {
            // Otherwise, create a new todo
            response = await fetch('/todos', {
                method: 'POST',
                body: formData,
            })

            if (!response.ok) {
                throw new Error('Failed to create todo')
            }

            todo = await response.json()

            // Add the new todo to the table
            const table = document.getElementById('todoTable')
            const row = table.insertRow()
            row.classList.add('todo-row')
            row.setAttribute('data-id', todo.id)
            row.insertCell().innerText = todo.title
            row.insertCell().innerText = todo.completed ? 'true' : 'false'
            const editCell = row.insertCell()
            const deleteCell = row.insertCell()
            editCell.innerHTML = '<button class="edit-button">Edit</button>'
            deleteCell.innerHTML =
                '<button class="delete-button">Delete</button>'
            bindButtonListeners(row)
        }

        // Clear the form fields after successful submission
        document.getElementById('title').value = ''
        document.getElementById('completed').checked = false
    } catch (error) {
        console.error('An error occurred:', error.message)
        // Here, you might want to show a user-friendly error message
    }
})

function bindButtonListeners(row) {
    const editButton = row.querySelector('.edit-button')
    const deleteButton = row.querySelector('.delete-button')

    editButton.addEventListener('click', (event) => {
        const row = event.target.parentElement.parentElement
        const id = row.getAttribute('data-id')
        const title = row.querySelector('.todo-title').innerText
        const completed =
            row.querySelector('.todo-completed').innerText === 'true'

        document.getElementById('title').value = title
        document.getElementById('completed').checked = completed

        form.setAttribute('data-id', id) // We'll use this when saving changes
    })

    deleteButton.addEventListener('click', async (event) => {
        const row = event.target.parentElement.parentElement
        const id = row.getAttribute('data-id')

        try {
            const response = await fetch(`/todos/${id}`, {
                method: 'DELETE',
            })

            if (!response.ok) {
                throw new Error('Failed to delete todo')
            }

            // If the deletion was successful, remove the row from the table
            row.remove()
        } catch (error) {
            console.error('An error occurred:', error.message)
            // Here, you might want to show a user-friendly error message
        }
    })
}

// Call bindButtonListeners for each existing row
document.querySelectorAll('.todo-row').forEach(bindButtonListeners)

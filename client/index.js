
const form = document.getElementById('todo__form')
const list = document.getElementById('todo__list')
const input = document.getElementById('todo__input')

const API_URL = "http://localhost:4000"
const DELETE_BUTTON_CLASS = 'todo__button_delete'

form.addEventListener('submit', async (event) => {
	event.preventDefault()

	await fetch(API_URL + "/todo/add", {
		method: 'POST',
		body: JSON.stringify(input.value)

	});

	await getTasks()
})

const getTasks = async () => {
	const response = await fetch(API_URL + "/tasks");
	const data = await response.json();

	clearList()
	setList(data)
}

const clearList = () => {
	list.innerHTML = ''
}

const setList = (data) => {
	data.forEach(({ Task, ID }) => {
		const p = document.createElement("li")

		p.innerHTML = Task
		list.appendChild(p)

		const button = document.createElement("button")
		button.className = DELETE_BUTTON_CLASS;
		button.type = 'button';
		button.dataset.ID = ID

		button.innerHTML = "удалить"
		p.appendChild(button)
	})

	addDeleteButtons()
}

const addDeleteButtons = () => {
	const deleteButtons = document.querySelectorAll(`.${DELETE_BUTTON_CLASS}`)

	Array.from(deleteButtons).forEach((element) => {
		element.addEventListener('click', () => deleteHandler(element))
	});
}

const deleteHandler = async (element) => {
	await fetch(API_URL + "/todo/delete", {
		method: 'POST',
		body: JSON.stringify(element.dataset.ID)
	})

	await getTasks()
}

getTasks()










document.addEventListener('DOMContentLoaded', () => {
    const todoForm = document.getElementById('todo-form');
    const todoInput = document.getElementById('todo-input');
    const todoList = document.getElementById('todo-list');
    const emptyState = document.getElementById('empty-state');

    // Загрузка задач из localStorage
    let todos = JSON.parse(localStorage.getItem('todos')) || [];
    renderTodos();

    // Обработчик добавления новой задачи
    todoForm.addEventListener('submit', (e) => {
        e.preventDefault();
        const text = todoInput.value.trim();
        
        if (text) {
            const newTodo = {
                id: Date.now(),
                text,
                completed: false
            };
            
            todos.push(newTodo);
            saveTodos();
            renderTodos();
            todoInput.value = '';
        }
    });

    // Обработчик изменения состояния задачи
    todoList.addEventListener('change', (e) => {
        if (e.target.classList.contains('checkbox')) {
            const id = parseInt(e.target.dataset.id);
            const todo = todos.find(todo => todo.id === id);
            if (todo) {
                todo.completed = e.target.checked;
                saveTodos();
                renderTodos();
            }
        }
    });

    // Обработчик удаления задачи
    todoList.addEventListener('click', (e) => {
        if (e.target.classList.contains('delete-btn')) {
            const id = parseInt(e.target.dataset.id);
            todos = todos.filter(todo => todo.id !== id);
            saveTodos();
            renderTodos();
        }
    });

    // Функция сохранения задач в localStorage
    function saveTodos() {
        localStorage.setItem('todos', JSON.stringify(todos));
    }

    // Функция отображения списка задач
    function renderTodos() {
        todoList.innerHTML = '';
        
        if (todos.length === 0) {
            emptyState.style.display = 'block';
            return;
        }
        
        emptyState.style.display = 'none';
        
        todos.forEach(todo => {
            const li = document.createElement('li');
            li.className = `todo-item ${todo.completed ? 'completed' : ''}`;
            
            li.innerHTML = `
                <input type="checkbox" class="checkbox" data-id="${todo.id}" ${todo.completed ? 'checked' : ''}>
                <span class="todo-text">${todo.text}</span>
                <div class="todo-actions">
                    <button class="delete-btn" data-id="${todo.id}">×</button>
                </div>
            `;
            
            todoList.appendChild(li);
        });
    }
}); 
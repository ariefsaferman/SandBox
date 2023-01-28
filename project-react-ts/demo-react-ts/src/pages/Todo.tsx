import React, { useState} from 'react'


interface ITodo {
    id: string;
    name: string;
    completed: boolean; 
}

export default function Todo(){
    const [inputTodo, setInputTodo] = useState<string>('');
    const [todos, setTodos] = useState<ITodo[]>([]);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>): void => {
        e.preventDefault(); 
        setInputTodo(e.target.value);
    }

    const handleSubmit = (e: React.FormEvent): void => {
        e.preventDefault();
        setTodos([...todos, {
            id: `${todos.length}-${inputTodo}`, 
            name: inputTodo, 
            completed: false
        }])

        setInputTodo(''); 
    }

    const handleDeleteTodo = (id: string): void => {
        const newTodos = todos.filter((todo) => todo.id !== id); 
        setTodos(newTodos); 
    }

    const handleCompleteTodo = (id: string): void => {
        const newTodos = todos.filter((todo) => {
            if (todo.id === id) {
                todo.completed = !todo.completed
            }
            return todo
        }); 
        setTodos(newTodos); 
    }

    return (
        <div>
            <h1>Todo</h1>
            <form onSubmit={(e) => handleSubmit(e)}>
                <input 
                type="text" 
                value={inputTodo} 
                onChange={(e) => handleChange(e)}
                placeholder="Input todo here.." 
                />
            </form>

            <h2>Todo List</h2>
            <ul>
                {
                    todos.map((todo: ITodo) => 
                        <li key={todo.id} className={todo.completed ? "completed" : ""}>{todo.name}
                            <button
                                className='delete-btn'
                                onClick={() => handleCompleteTodo(todo.id)}>
                                    ✅
                            </button>

                            <button
                                className='delete-btn'
                                onClick={() => handleDeleteTodo(todo.id)}>🗑
                            </button>

                        </li>

                    )
                }
            </ul>
        </div>
    )
}
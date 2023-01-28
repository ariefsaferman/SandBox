import React, { useState } from "react";

interface ITodo {
	id: string;
	name: string;
	completed: boolean;
}

export default function Todo() {
	const [todos, setTodos] = useState<ITodo[]>([]);
	const [inputTodo, setInputTodo] = useState<string>("");

	const handleChangeTodo = (
		e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
	): void => {
		setInputTodo(e.target.value);
	};

	const handleAddTodo = (e: React.FormEvent): void => {
		e.preventDefault();
		setTodos([
			...todos,
			{
				id: `${todos.length}-${inputTodo}`,
				name: inputTodo,
				completed: false,
			},
		]);

		setInputTodo("");
	};

	const handleCompleteTodo = (id: string): void => {
		const newTodos = todos.filter((todo) => {
			if (todo.id === id) {
				todo.completed = !todo.completed;
			}
			return todo;
		});

		setTodos(newTodos);
	};

	const handleDeleteTodo = (id: string): void => {
		const newTodos = todos.filter((todo) => todo.id !== id);

		setTodos(newTodos);
	};

	return (
		<div>
			<h1>Todo</h1>
			<form onSubmit={handleAddTodo}>
				<input
					type="text"
					value={inputTodo}
					onChange={handleChangeTodo}
					placeholder="Input todo here.."
				/>
			</form>
			<h2>Todo Lists</h2>
			<ul>
				{todos.map((todo) => (
					<li key={todo.id} className={todo.completed ? "completed" : ""}>
						{todo.name}
						<button
							className="delete-btn"
							onClick={() => handleCompleteTodo(todo.id)}
						>
							{!todo.completed ? "✅" : "↩"}
						</button>
						<button
							className="delete-btn"
							onClick={() => handleDeleteTodo(todo.id)}
						>
							🗑
						</button>
					</li>
				))}
			</ul>
		</div>
	);
}

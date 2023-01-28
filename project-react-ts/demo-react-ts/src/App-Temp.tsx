import React, { useState } from 'react';
import Todo from './pages/Todo';

import './App.css';

function App() {
  const [name, setName] = useState<string>("");
  const [arr, setArr] = useState<string[]>([]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setName(e.target.value); 
  }

  const clickBtn = (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault(); 
    setArr([...arr, name])
  }

  return (
    <div className="App">
      <p>Hello {name}</p>
      <form action="">
        <input type="text" onChange={(e) => handleChange(e)} />
        <button onClick={(e) => clickBtn(e)}>Add Name</button>
      </form>
      <ul>
        {
          arr.map((item: string, index: number) => {
            return <li key={index}>{item}</li>
        })
        }
      </ul>
    </div>
  );
}

export default App;

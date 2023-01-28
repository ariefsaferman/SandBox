import React, { useState } from 'react';
import Todo from './pages/Todo';
import AppEffect from './App-useEffect';

import './App.css';
import Posts from './pages/Posts';
import PostCustom from './pages/PostCustom';

function App() {
  return (
    <div className='App'>
      {/* <Todo></Todo> */}
      {/* <AppEffect/> */}
      {/* <Posts/> */}
      <PostCustom></PostCustom>
    </div>
  );
}

export default App;

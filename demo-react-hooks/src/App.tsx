import Todo from "./Pages/Todo";
import Posts from "./Pages/Posts";
import PostsCustom from "./Pages/PostsCustom";
import Layout from "./component/Layout";
import { Routes, Route } from "react-router-dom";
import "./App.css";
import Profile from "./Pages/Profile";
import ProtectedPage from "./Pages/ProtectedPage";
import Login from "./Pages/Login";
import React from "react";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="" element={<Layout />}>
          <Route index element={<Todo />} />
          <Route path="posts">
            <Route index element={<Posts />} />
            <Route path="custom" element={<PostsCustom />} />
          </Route>
          <Route element={<ProtectedPage />}>
            <Route path="profile" element={<Profile />} />
          </Route>
        </Route>
        <Route path="login" element={<Login />} />
        <Route path="" element={<h2> 404 not found!</h2>}></Route>
      </Routes>
    </div>
  );
}

export default App;

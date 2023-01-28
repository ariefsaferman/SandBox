import { Route, Routes } from "react-router-dom";
import AddBookPage from "./pages/AddBook";
import EditBookPage from "./pages/EditBook";
import LoginPage from "./pages/Login";
import RegisterPage from "./pages/Register";
import ProtectedPage from "./pages/ProtectedPage";
import Home from "./pages/Home";

import "bootstrap/dist/css/bootstrap.min.css";
import "./App.css";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route element={<ProtectedPage />}>
          <Route path="/" element={<Home />} />
          <Route path="/add" element={<AddBookPage />} />
          <Route path="/edit/:id" element={<EditBookPage />} />
        </Route>
        <Route path="/login" element={<LoginPage />}></Route>
        <Route path="/register" element={<RegisterPage />}></Route>
        <Route path="*" element={<h2>404 not found!</h2>}></Route>
      </Routes>
    </div>
  );
}

export default App;

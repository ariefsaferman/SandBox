import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { RegisterRequest } from "../interface/register";

export default function RegisterPage() {
  const [email, setEmail] = useState("");
  const [userName, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleRegisterClick = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const registerRequest: RegisterRequest = {
      userName: userName,
      email: email,
      password: password,
    };

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(registerRequest),
    };

    fetch("http://localhost:3004/register", requestOptions)
      .then((res) => {
        if (res.ok) {
          navigate("/");
        }
        throw new Error("Error register");
      })
      .catch((err) => console.log(err));
  };

  return (
    <div className="d-flex vh-100 flex-column align-items-center justify-content-center">
      <div>
        <h1 className="text-start fw-bold font__size">
          The <br />
          Library <br />
          Books
        </h1>
        <form onClick={handleRegisterClick}>
          <div className="input-group mb-3">
            <input
              type="text"
              className="form-control control_login"
              placeholder="Username"
              aria-label="Username"
              aria-describedby="basic-addon1"
              onChange={(e) => {
                e.preventDefault();
                setUserName(e.target.value);
              }}
              required
            />
          </div>

          <div className="input-group mb-3">
            <input
              type="email"
              className="form-control control_login"
              placeholder="Email"
              aria-label="Username"
              aria-describedby="basic-addon1"
              onChange={(e) => {
                e.preventDefault();
                setEmail(e.target.value);
              }}
              required
            />
          </div>

          <div className="input-group mb-3">
            <input
              type="password"
              className="form-control control_login"
              placeholder="password"
              aria-label="Username"
              aria-describedby="basic-addon1"
              onChange={(e) => {
                e.preventDefault();
                setPassword(e.target.value);
              }}
              required
            />
          </div>

          <div className="input-group mb-3">
            <input
              type="password"
              className="form-control control_login"
              placeholder="Confirm Password"
              aria-label="Username"
              aria-describedby="basic-addon1"
              onChange={(e) => {
                e.preventDefault();
              }}
              required
            />
          </div>
          <div className="input-group mb-3 justify-content-center">
            <button
              type="submit"
              className="btn btn-info btn-lg w-100 text-white"
            >
              Register
            </button>
          </div>

          <div>
            <p>
              Already have an account{" "}
              <Link to="/login" className="text-decoration-none text-black">
                Login
              </Link>
            </p>
          </div>
        </form>
      </div>

      <div id="footer">Build for learning purpose</div>
    </div>
  );
}

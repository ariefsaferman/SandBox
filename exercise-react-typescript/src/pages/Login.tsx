import React, { useState } from "react";
import { Link, useLocation, useNavigate } from "react-router-dom";
import AuthConsumer from "../hooks/useAuth";
import { LoginRequest } from "../interface/login";
import dotenv from "dotenv-webpack";



export default function LoginPage() {
  const API_LOGIN: string = process.env.API_LOGIN || "";
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const authState = AuthConsumer();
  const navigate = useNavigate();
  const location = useLocation();
  let from = location.state?.from?.pathname || "/";

  const handleLoginClick = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const loginRequest: LoginRequest = {
      email: email,
      password: password,
    };

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(loginRequest),
    };

    fetch(API_LOGIN, requestOptions)
      .then((props) => {
        if (props.ok) {
          return props.json();
        }
        throw new Error("error");
      })
      .then((props) => {
        authState.setAuthenticated?.(true);
        localStorage.setItem("token", JSON.stringify(props.accessToken));
        navigate(from, { replace: true });
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
        <form onClick={handleLoginClick}>
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
          <div className="input-group mb-3 justify-content-center">
            <button
              type="submit"
              className="btn btn-info btn-lg w-100 text-white"
            >
              Login
            </button>
          </div>

          <div>
            <p>
              You don't have an account?,{"  "}
              <Link to="/register" className="text-decoration-none text-black">
                Signup
              </Link>
            </p>
          </div>
        </form>
      </div>

      <div id="footer">Build for learning purpose</div>
    </div>
  );
}

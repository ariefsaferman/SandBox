import React from "react";
import { useNavigate } from "react-router-dom";
import AuthConsumer from "../hooks/useAuth";

export default function Profile() {
  const authState = AuthConsumer();
  const navigate = useNavigate();

  const handleClick = () => {
    authState.setAuthenticated?.(false);


    navigate("/", {replace: true});
  };

  return (
    <div>
      <h1>Profile</h1>
      <button onClick={handleClick}>Log Out</button>
    </div>
  );
}

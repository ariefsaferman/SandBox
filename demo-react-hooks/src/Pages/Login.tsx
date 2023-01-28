import { useNavigate, useLocation } from "react-router-dom";
import AuthConsumer from "../hooks/useAuth";

export default function Login() {
  const authState = AuthConsumer();
  const navigate = useNavigate();
  const location = useLocation();

  // get previous page from useLocation
  //   console.log('login', location);
  let from = location.state?.from?.pathname || "/";

  const handleClick = () => {
    authState.setAuthenticated?.(true);
    // authState.setAuthenticated!(true)

    // replace previous history so the user cannot go to login when clicking back button in the browser
    navigate(from, { replace: true });
  };

  return (
    <div>
      <h1>Login</h1>
      <button onClick={handleClick}>Login</button>
    </div>
  );
}

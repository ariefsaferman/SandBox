import AuthConsumer from "../hooks/useAuth";
import { Navigate, Outlet, useLocation } from "react-router-dom";

export default function ProtectedPage() {
  const authState = AuthConsumer();
  let location = useLocation();

  // if not authenticated then redirect to the login page
  if (!authState.isAuth) {
    return <Navigate to="login" replace state={{ from: location }} />;
    //TODO: add location and replace from react router
  }

  return <Outlet />;
}

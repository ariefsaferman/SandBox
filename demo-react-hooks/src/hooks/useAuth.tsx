import React, { createContext, Dispatch, useState, useContext } from "react";

// define the context tpe
type AuthType = {
  isAuth: boolean;
  setAuthenticated?: Dispatch<React.SetStateAction<boolean>>;
};

// create the initial state
const initialState: AuthType = { isAuth: false };

// create the context
const AuthContext = createContext<AuthType>(initialState);

// create Provider
export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuth, setAuthenticated] = useState<boolean>(false);

  return (
    <AuthContext.Provider value={{ isAuth, setAuthenticated }}>
      {children}
    </AuthContext.Provider>
  );
}

// create Consumer
export default function AuthConsumer() {
  return useContext(AuthContext);
}



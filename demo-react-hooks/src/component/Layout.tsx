import React from "react";
import { Link, Outlet, NavLink } from "react-router-dom";

export default function Layout() {
  return (
    <div>
      <nav>
        <ul className="navbar">
          <li>
            <NavLink to="/">Todo</NavLink>
          </li>
          <li>
            <NavLink to="/posts">Posts</NavLink>
          </li>
          <li>
            <Link to="/posts/custom">Post Custom</Link>
          </li>
          <li>
            <Link to="/profile">Profile</Link>
          </li>
        </ul>
      </nav>
      <hr />
      <Outlet />
    </div>
  );
}

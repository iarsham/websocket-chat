import React from "react";
import "./Header.css";

const Header: React.FC = () => {
  return (
    <header className="header">
      <h1>Websocket Chat APP</h1>
      <nav>
        <a href="/client/public">SignUp</a>
        <a href="/login">Login</a>
      </nav>
    </header>
  );
};

export default Header;

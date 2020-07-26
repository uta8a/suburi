import React from "react";
import { CTF_NAME } from "../../util/global";
import { Link } from "react-router-dom";

const Header = () => {
  const headerComponentStyle = "text-gray-100 font-light mx-4 hover:underline";
  return (
    <header>
      <nav className="bg-gray-800 shadow-sm">
        <div className="container flex justify-between items-center mx-auto px-6 py-4">
          <div>
            <Link to="/" className="text-xl text-white">
              {CTF_NAME}
            </Link>
          </div>
          <div>
            <Link to="/notifications" className={headerComponentStyle}>
              Notifications
            </Link>
            <Link to="/about" className={headerComponentStyle}>
              About
            </Link>
            <Link to="/rules" className={headerComponentStyle}>
              Rules
            </Link>
            <Link to="/scoreboard" className={headerComponentStyle}>
              Scoreboard
            </Link>
            <Link to="/challenge" className={headerComponentStyle}>
              Challenge
            </Link>
            {/* TODO islogin ? hidden : show */}
            <Link to="/login" className={headerComponentStyle}>
              Login
            </Link>
            <Link to="/register" className={headerComponentStyle}>
              Register
            </Link>
          </div>
        </div>
      </nav>
    </header>
  );
};

export { Header };

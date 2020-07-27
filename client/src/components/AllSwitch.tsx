import React, { useState, useEffect, Props } from "react";
import { Switch, Route, useLocation } from "react-router-dom";
import { Login } from "./Login";
import { Home } from "./Home";
import { Notifications } from "./Notifications";
import { About } from "./About";
import { Rules } from "./Rules";
import { Scoreboard } from "./Scoreboard";
import { Challenge } from "./Challenge";
import { Register } from "./Register";
import { CTF_NAME } from "../util/global";
const pathToTitle = (path: string) => {
  return path.charAt(1).toUpperCase() + path.slice(2);
};
const AllSwitch = () => {
  const noScrollPath = ["/", "/about", "/rules", "/login", "/register"];
  let location = useLocation();
  let mainStyle = "flex-grow h-full";
  useEffect(() => {
    document.title = `${pathToTitle(location.pathname)} ${
      location.pathname === "/" ? "" : "|"
    } ${CTF_NAME}`;
  }, [location]);
  return (
    <main
      className={`${mainStyle} ${
        noScrollPath.includes(location.pathname)
          ? "overflow-y-hidden"
          : "overflow-y-scroll"
      }`}
    >
      <Switch>
        <Route exact path="/" component={Home} />
        <Route exact path="/user-check" component={Page} />
        <Route exact path="/tester-check" component={Page} />
        <Route exact path="/admin-check" component={Page} />
        <Route exact path="/get-token" component={Page} /> {/* for debug */}
        <Route exact path="/login" component={Login} />
        <Route exact path="/register" component={Register} />
        <Route exact path="/notifications" component={Notifications} />
        <Route exact path="/about" component={About} />
        <Route exact path="/rules" component={Rules} />
        <Route exact path="/scoreboard" component={Scoreboard} />
        <Route exact path="/challenge" component={Challenge} />
      </Switch>
    </main>
  );
};

const Page = () => {
  return <h1 className="text-3xl">Hello~!</h1>;
};
export { AllSwitch };

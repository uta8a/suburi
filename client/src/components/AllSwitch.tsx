import React from "react";
import { Switch, Route } from "react-router-dom";
import { Login } from "./Login";
import { Home } from "./Home";
import { Notifications } from "./Notifications";
import { About } from "./About";
import { Rules } from "./Rules";
import { Scoreboard } from "./Scoreboard";
import { Challenge } from "./Challenge";
import { Register } from "./Register";

const AllSwitch = () => {
  return (
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
  );
};

const Page = () => {
  return <h1 className="text-3xl">Hello~!</h1>;
};
export { AllSwitch };

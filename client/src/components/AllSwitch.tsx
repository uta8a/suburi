import React from "react";
import { Switch, Route } from "react-router-dom";
import { Login } from "./Login";
import { Home } from "./Home";
import { Notification } from "./Notification";
import { About } from "./About";

const AllSwitch = () => {
  return (
    <Switch>
      <Route exact path="/" component={Home} />
      <Route exact path="/user-check" component={Page} />
      <Route exact path="/tester-check" component={Page} />
      <Route exact path="/admin-check" component={Page} />
      <Route exact path="/get-token" component={Page} /> {/* for debug */}
      <Route exact path="/login" component={Login} />
      <Route exact path="/register" component={Page} />
      <Route exact path="/notification" component={Notification} />
      <Route exact path="/about" component={About} />
    </Switch>
  );
};

const Page = () => {
  return <h1 className="text-3xl">Hello~!</h1>;
};
export { AllSwitch };

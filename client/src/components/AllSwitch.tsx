import React from "react";
import { Switch, Route } from "react-router-dom";
import { Login } from "./Login";

const AllSwitch = () => {
  return (
    <Switch>
      <Route exact path="/" component={Page} />
      <Route exact path="/user-check" component={Page} />
      <Route exact path="/tester-check" component={Page} />
      <Route exact path="/admin-check" component={Page} />
      <Route exact path="/get-token" component={Page} /> {/* for debug */}
      <Route exact path="/login" component={Login} />
      <Route exact path="/register" component={Page} />
    </Switch>
  );
};

const Page = () => {
  return <h1 className="text-3xl">Hello~!</h1>;
};
export { AllSwitch };

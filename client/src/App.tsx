import React, { useState, useEffect } from "react";
import { Request } from "./proto/user/user_pb";
import { UserClient } from "./proto/user/UserServiceClientPb";
import useSWR from "swr";
import { URL } from "./util/global";
import { Login } from "./components/Login";
import { AllRoute } from "./components/AllRoute";
import { AllSwitch } from "./components/AllSwitch";
import { BrowserRouter as Router, useLocation } from "react-router-dom";
import { Header } from "./components/view/Header";
import { Footer } from "./components/view/Footer";

interface ioState {
  // input
  username: string;
  password: string;
  // output
  message: string;
}

const App = () => {
  const [{ token, message }, setState] = useState({ token: "", message: "" });
  const onChange = () => {
    const request = new Request();
    request.setUsername("");
    request.setPassword("");
    const client = new UserClient(`${URL}`, {}, {});
    client
      .getToken(request, null)
      .then((ret) => {
        setState({ token: ret.getToken(), message: ret.getDisplayMessage() });
      })
      .catch((err) => {
        throw err;
      });
  };
  const fetcher = () => {
    const request = new Request();
    request.setUsername("");
    request.setPassword("");
    const client = new UserClient(`${URL}`, {}, {});
    const ret = client.getToken(request, null).catch((err) => {
      throw err;
    });
    return ret;
  };
  const { data } = useSWR("/getToken", fetcher);

  const appStyle = "flex flex-col h-screen bg-gray-900 overflow-hidden";
  return (
    <div className={appStyle}>
      <Router>
        <Header />
        <AllSwitch />
        <Footer />
      </Router>
    </div>
  );
};

export default App;

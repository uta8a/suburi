import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';

import { HelloRequest } from "./helloworld/helloworld_pb";
import { GreeterClient } from "./helloworld/HelloworldServiceClientPb";

const URL = "http://localhost:8080"

const initialState = {
  inputText: "World",
  message: ""
};

const App = () => {
  const [
    { inputText, message },
    setState
  ] = useState(initialState);

  const onClick = () => {
    const request = new HelloRequest();
    request.setName(inputText);

    const client = new GreeterClient(`${URL}`, {}, {});
    client.sayHello(request, {}, (err, ret) => {
      if (err || ret == null) {
        throw err;
      }
      setState({ message: ret.getMessage() });
    })
  }
  const onChange = ( e:React.ChangeEvent<HTMLInputElement> ) => {
    setState({ inputText: e.target.value })
  }
  return (
    <div className="App">
      <input
        type="text"
        value={inputText}
        onChange={onChange}
      />
      <button onClick={onClick}>Send</button>
      <p>{message}</p>
    </div>
  )
}

export default App;

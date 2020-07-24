import React, { useState } from 'react';
import { Request } from './proto/user/user_pb';
import { UserClient } from './proto/user/UserServiceClientPb';

// TODO normal gRPC

const URL = 'http://localhost:8080'

interface ioState {
  // input
  username: string;
  password: string;
  // output
  message: string;
}

const App = () => {
  const [
    { token, message },
    setState,
  ] = useState({ token: "", message: ""})
  const onChange = () => {
    const request = new Request();
    request.setUsername("");
    request.setPassword("");
    const client = new UserClient(`${URL}`, {}, {});
    client.getToken(request, null).then(ret => {
      setState({ token: ret.getToken(), message: ret.getDisplayMessage() })
    }).catch(err => {
      throw err;
    });
  }
  return (
    <div className="App">
      {/* grpc */}
      <button onClick={onChange}>
        Get Token!!
      </button>
      <h1>{token}</h1>
      <h2>{message}</h2>
      <p>{11111}</p>
    </div>
  );
}

export default App;

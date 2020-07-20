import React, { useState } from 'react';
import logo from './logo.svg';
import './App.css';

// import { TellMeSecretRequest } from "./helloworld/helloworld_pb";
// import { GreeterClient } from "./helloworld/HelloworldServiceClientPb";

const URL = "http://localhost:8080"

interface MessageState {
  inputText: string;
  message: string;
}

const initialState = {
  inputText: "alice",
  message: ""
};

const App = () => {
  const [
    { inputText, message },
    setState
  ] = useState<MessageState>(initialState);

  const onClick = () => {
    //const request = new TellMeSecretRequest();
    //request.setMessage(inputText);
    // const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhbGljZSIsIm5hbWUiOiJKb2huIERvZSIsImlhdCI6MTUxNjIzOTAyMn0.vqbJV8ftQ1Tx_UMqUioMd5jEve9Odpw9--NJLsQOSeo';
    const token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJib2IiLCJpYXQiOjE1MTYyMzkwMjJ9.YXGHLm79VdQH5Ccre7MxKXHTcbDMb4sz2uYqiQE21Vg';
    const metadata = {'authorization': `bearer ${token}`};
    //const client = new GreeterClient(`${URL}`, {}, {});
    //  client.tellMeSecret(request, metadata, (err, ret) => {
    //    if (err || ret == null) {
    //      throw err;
    //    }
    //    setState({ inputText: inputText, message: ret.getAnswer() });
    //  })
  }
  const onChange = ( e:React.ChangeEvent<HTMLInputElement> ) => {
    setState({ inputText: e.target.value, message: message })
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

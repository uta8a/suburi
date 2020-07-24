import React, { useState } from 'react';
// TODO normal gRPC
interface ioState {
  // input
  username: string;
  password: string;
  // output
  message: string;
}
const initialState = {
  username: "",
  password: "",
  message: "",
}

const App = () => {
  const [
    { username, password, message },
    setState,
  ] = useState<ioState>(initialState);
  const 
  return (
    <div className="App">
      {/* grpc */}
    </div>
  );
}

export default App;

import React, { FunctionComponent } from "react";

const Title: FunctionComponent = (props) => {
  return (
    <h1 className="text-center pt-6 font-bold text-white text-3xl">
      {props.children}
    </h1>
  );
};

export { Title };

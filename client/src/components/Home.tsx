import React from "react";
import { CTF_NAME, CTF_SUBTITLE } from "../util/global";
import '../fonts.css'


const Home = () => {
  return (
    <div className="flex w-full h-full flex-col justify-center items-center">
      <h1 className="my-6 text-center text-white text-5xl font-bold custom-font-family">
        {CTF_NAME}
      </h1>
      <p className="my-6  text-white text-center">
        <CTF_SUBTITLE />
      </p>
      <div className="ctf-timer my-6 text-white text-center text-4xl">
        0:00:00
      </div>
    </div>
  );
};

export { Home };

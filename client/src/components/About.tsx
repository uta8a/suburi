import React from "react";
import { CTF_NAME, CTF_ADMIN_NAME, CTF_ADMIN_LINK } from "../util/global";
import { Title } from "./view/Title";

const Description = () => {
  const staffStyle = "text-center pt-3 text-white text-xl";
  return (
    <>
      {/* period */}
      <div className="flex flex-col pt-3">
        <h2 className="text-center pt-3 text-white text-2xl">Period</h2>
        <p className="text-center pt-3 text-white text-3xl">
          01/01 00:00 - 01/03 00:00
        </p>
      </div>
      {/* staff */}
      <div className="flex flex-col pt-3">
        <h2 className="text-center pt-3 text-white text-2xl">Staff</h2>
        <p className={staffStyle}>
          {CTF_NAME} is organized by{" "}
          <a className="hover:underline" href={CTF_ADMIN_LINK}>
            {CTF_ADMIN_NAME}
          </a>
          .
        </p>
        <p className={staffStyle}>頑張って作ったのでぜひ楽しんでください〜</p>
      </div>
    </>
  );
};

const About = () => {
  return (
    <>
      <Title>About</Title>
      <Description />
    </>
  );
};

export { About };

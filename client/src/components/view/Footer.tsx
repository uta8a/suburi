import React from "react";
import { CTF_NAME, CTF_ADMIN_NAME, CTF_ADMIN_LINK } from "../../util/global";

const Footer = () => {
  return (
    <footer className="bg-gray-800">
      <p className="p-2 text-white text-center text-base">
        {CTF_NAME} is made by{" "}
        <a
          className="hover:underline"
          href={CTF_ADMIN_LINK}
          target="_blank"
          rel="noopener noreferrer"
        >
          {CTF_ADMIN_NAME}
        </a>
        .
      </p>
    </footer>
  );
};
export { Footer };

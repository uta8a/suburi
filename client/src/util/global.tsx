import React from "react";
const URL = "http://localhost:8080";
const CTF_NAME = "ErilyCTF";
const CTF_ADMIN_NAME = "@kaito_tateyama";
const CTF_ADMIN_LINK = "https://twitter.com/kaito_tateyama";
const CTF_SUBTITLE = () => {
  return (
    <>
      🚩🚩🚩 これは試験的に行われます。フィードバックは
      <a className="bg-green-500 rounded-md underline" href="#">
        こちら
      </a>
      からお願いします。 🚩🚩🚩
    </>
  );
};

export { URL, CTF_NAME, CTF_ADMIN_NAME, CTF_ADMIN_LINK, CTF_SUBTITLE };
